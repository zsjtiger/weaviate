//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package grouper

import (
	"fmt"
	"strings"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/search"
)

type valueType int

const (
	numerical valueType = iota
	textual
	boolean
	reference
	geo
	unknown
)

type valueGroup struct {
	values    []interface{}
	valueType valueType
	name      string
}

func (g group) flattenMerge() (search.Result, error) {
	values := g.makeValueGroups()
	merged, err := mergeValueGroups(values)
	if err != nil {
		return search.Result{}, fmt.Errorf("merge values: %v", err)
	}

	vector, err := g.mergeVectors()
	if err != nil {
		return search.Result{}, fmt.Errorf("merge vectors: %v", err)
	}

	return search.Result{
		Schema: merged,
		Vector: vector,
	}, nil
}

func (g group) makeValueGroups() map[string]valueGroup {
	values := map[string]valueGroup{}

	for _, elem := range g.Elements {
		if elem.Schema == nil {
			continue
		}

		for propName, propValue := range elem.Schema.(map[string]interface{}) {
			current, ok := values[propName]
			if !ok {
				current = valueGroup{
					values:    []interface{}{propValue},
					valueType: valueTypeOf(propValue),
					name:      propName,
				}
				values[propName] = current
				continue
			}

			current.values = append(current.values, propValue)
			values[propName] = current
		}
	}

	return values
}

func (g group) mergeVectors() ([]float32, error) {
	amount := len(g.Elements)
	if amount == 0 {
		return nil, nil
	}

	if amount == 1 {
		return g.Elements[0].Vector, nil
	}

	dimensions := len(g.Elements[0].Vector)
	out := make([]float32, dimensions, dimensions)

	// sum up
	for _, groupElement := range g.Elements {
		if len(groupElement.Vector) != dimensions {
			return nil, fmt.Errorf("vectors have different dimensions")
		}

		for i, vectorElement := range groupElement.Vector {
			out[i] = out[i] + vectorElement
		}
	}

	// divide by amount of vectors
	for i := range out {
		out[i] = out[i] / float32(amount)
	}

	return out, nil
}

func mergeValueGroups(props map[string]valueGroup) (map[string]interface{}, error) {
	mergedProps := map[string]interface{}{}

	for propName, group := range props {
		var (
			res interface{}
			err error
		)
		switch group.valueType {
		case textual:
			res, err = mergeTextualProps(group.values)
		case numerical:
			res, err = mergeNumericalProps(group.values)
		case boolean:
			res, err = mergeBooleanProps(group.values)
		case geo:
			res, err = mergeGeoProps(group.values)
		case reference:
			res, err = mergeReferenceProps(group.values)
		case unknown:
			continue
		default:
			err = fmt.Errorf("unrecognized value type")
		}
		if err != nil {
			return nil, fmt.Errorf("prop '%s': %v", propName, err)
		}

		mergedProps[propName] = res
	}

	return mergedProps, nil
}

func valueTypeOf(in interface{}) valueType {
	switch in.(type) {
	case string:
		return textual
	case float64:
		return numerical
	case bool:
		return boolean
	case *models.GeoCoordinates:
		return geo
	case []interface{}:
		return reference
	default:
		return unknown
	}
}

func mergeTextualProps(in []interface{}) (string, error) {
	var values []string
	seen := make(map[string]struct{}, len(in))
	for i, elem := range in {
		asString, ok := elem.(string)
		if !ok {
			return "", fmt.Errorf("element %d: expected textual element to be string, but got %T", i, elem)
		}

		if _, ok := seen[asString]; ok {
			// this is a duplicate, don't append it again
			continue
		}

		seen[asString] = struct{}{}
		values = append(values, asString)
	}

	if len(values) == 1 {
		return values[0], nil
	}

	return fmt.Sprintf("%s (%s)", values[0], strings.Join(values[1:], ", ")), nil
}

func mergeNumericalProps(in []interface{}) (float64, error) {
	var sum float64
	for i, elem := range in {
		asFloat, ok := elem.(float64)
		if !ok {
			return 0, fmt.Errorf("element %d: expected numerical element to be float64, but got %T", i, elem)
		}

		sum += asFloat
	}

	return sum / float64(len(in)), nil
}

func mergeBooleanProps(in []interface{}) (bool, error) {
	var countTrue uint
	var countFalse uint
	for i, elem := range in {
		asBool, ok := elem.(bool)
		if !ok {
			return false, fmt.Errorf("element %d: expected boolean element to be bool, but got %T", i, elem)
		}

		if asBool {
			countTrue++
		} else {
			countFalse++
		}
	}

	return countTrue >= countFalse, nil
}

func mergeGeoProps(in []interface{}) (*models.GeoCoordinates, error) {
	var sumLat float32
	var sumLon float32

	for i, elem := range in {
		asGeo, ok := elem.(*models.GeoCoordinates)
		if !ok {
			return nil, fmt.Errorf("element %d: expected geo element to be *models.GeoCoordinates, but got %T", i, elem)
		}

		if asGeo.Latitude != nil {
			sumLat += *asGeo.Latitude
		}
		if asGeo.Longitude != nil {
			sumLon += *asGeo.Longitude
		}
	}

	return &models.GeoCoordinates{
		Latitude:  ptFloat32(sumLat / float32(len(in))),
		Longitude: ptFloat32(sumLon / float32(len(in))),
	}, nil
}

func ptFloat32(in float32) *float32 {
	return &in
}

func mergeReferenceProps(in []interface{}) ([]interface{}, error) {
	var out []interface{}
	seenID := map[string]struct{}{}

	for i, elem := range in {
		asSlice, ok := elem.([]interface{})
		if !ok {
			return nil, fmt.Errorf("element %d: expected reference values to be slice, but got %T", i, elem)
		}

		for _, singleRef := range asSlice {
			asRef, ok := singleRef.(search.LocalRef)
			if !ok {
				// don't know what to do with this type, ignore
				continue
			}

			id, ok := asRef.Fields["uuid"]
			if !ok {
				return nil, fmt.Errorf("found a search.LocalRef, but 'uuid' field is missing: %#v", asRef)
			}

			if _, ok := seenID[id.(string)]; ok {
				// duplicate
				continue
			}

			out = append(out, asRef)
			seenID[id.(string)] = struct{}{} // make sure we skip this next time
		}

	}

	return out, nil
}