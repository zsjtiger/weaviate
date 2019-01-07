/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateSchemaActionsDeleteHandlerFunc turns a function with the right signature into a weaviate schema actions delete handler
type WeaviateSchemaActionsDeleteHandlerFunc func(WeaviateSchemaActionsDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateSchemaActionsDeleteHandlerFunc) Handle(params WeaviateSchemaActionsDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateSchemaActionsDeleteHandler interface for that can handle valid weaviate schema actions delete params
type WeaviateSchemaActionsDeleteHandler interface {
	Handle(WeaviateSchemaActionsDeleteParams, interface{}) middleware.Responder
}

// NewWeaviateSchemaActionsDelete creates a new http.Handler for the weaviate schema actions delete operation
func NewWeaviateSchemaActionsDelete(ctx *middleware.Context, handler WeaviateSchemaActionsDeleteHandler) *WeaviateSchemaActionsDelete {
	return &WeaviateSchemaActionsDelete{Context: ctx, Handler: handler}
}

/*WeaviateSchemaActionsDelete swagger:route DELETE /schema/actions/{className} schema weaviateSchemaActionsDelete

Remove an Action class (and all data in the instances) from the ontology.

*/
type WeaviateSchemaActionsDelete struct {
	Context *middleware.Context
	Handler WeaviateSchemaActionsDeleteHandler
}

func (o *WeaviateSchemaActionsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateSchemaActionsDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
