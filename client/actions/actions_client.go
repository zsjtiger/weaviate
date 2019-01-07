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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new actions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for actions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WeaviateActionHistoryGet gets an action s history based on its UUID related to this key

Returns a particular Action history.
*/
func (a *Client) WeaviateActionHistoryGet(params *WeaviateActionHistoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionHistoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionHistoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.action.history.get",
		Method:             "GET",
		PathPattern:        "/actions/{actionId}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionHistoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionHistoryGetOK), nil

}

/*
WeaviateActionUpdate updates an action based on its UUID related to this key

Updates an Action's data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) WeaviateActionUpdate(params *WeaviateActionUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionUpdateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.action.update",
		Method:             "PUT",
		PathPattern:        "/actions/{actionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionUpdateAccepted), nil

}

/*
WeaviateActionsCreate creates actions between two things object and subject

Registers a new Action. Provided meta-data and schema values are validated.
*/
func (a *Client) WeaviateActionsCreate(params *WeaviateActionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsCreateOK, *WeaviateActionsCreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.create",
		Method:             "POST",
		PathPattern:        "/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *WeaviateActionsCreateOK:
		return value, nil, nil
	case *WeaviateActionsCreateAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
WeaviateActionsDelete deletes an action based on its UUID related to this key

Deletes an Action from the system.
*/
func (a *Client) WeaviateActionsDelete(params *WeaviateActionsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.delete",
		Method:             "DELETE",
		PathPattern:        "/actions/{actionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsDeleteNoContent), nil

}

/*
WeaviateActionsGet gets a specific action based on its UUID and a thing UUID related to this key also available as websocket bus

Lists Actions.
*/
func (a *Client) WeaviateActionsGet(params *WeaviateActionsGetParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.get",
		Method:             "GET",
		PathPattern:        "/actions/{actionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsGetOK), nil

}

/*
WeaviateActionsList gets a list of actions related to this key

Lists all Actions in reverse order of creation, owned by the user that belongs to the used token.
*/
func (a *Client) WeaviateActionsList(params *WeaviateActionsListParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.list",
		Method:             "GET",
		PathPattern:        "/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsListOK), nil

}

/*
WeaviateActionsPatch updates an action based on its UUID using patch semantics related to this key

Updates an Action. This method supports patch semantics. Provided meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) WeaviateActionsPatch(params *WeaviateActionsPatchParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsPatchOK, *WeaviateActionsPatchAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.patch",
		Method:             "PATCH",
		PathPattern:        "/actions/{actionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *WeaviateActionsPatchOK:
		return value, nil, nil
	case *WeaviateActionsPatchAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
WeaviateActionsPropertiesCreate adds a single reference to a class property when cardinality is set to has many

Add a single reference to a class-property when cardinality is set to 'hasMany'.
*/
func (a *Client) WeaviateActionsPropertiesCreate(params *WeaviateActionsPropertiesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsPropertiesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsPropertiesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.properties.create",
		Method:             "POST",
		PathPattern:        "/actions/{actionId}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsPropertiesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsPropertiesCreateOK), nil

}

/*
WeaviateActionsPropertiesDelete deletes the single reference that is given in the body from the list of references that this property has

Delete the single reference that is given in the body from the list of references that this property has.
*/
func (a *Client) WeaviateActionsPropertiesDelete(params *WeaviateActionsPropertiesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsPropertiesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsPropertiesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.properties.delete",
		Method:             "DELETE",
		PathPattern:        "/actions/{actionId}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsPropertiesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsPropertiesDeleteNoContent), nil

}

/*
WeaviateActionsPropertiesUpdate replaces all references to a class property

Replace all references to a class-property.
*/
func (a *Client) WeaviateActionsPropertiesUpdate(params *WeaviateActionsPropertiesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsPropertiesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsPropertiesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.properties.update",
		Method:             "PUT",
		PathPattern:        "/actions/{actionId}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsPropertiesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsPropertiesUpdateOK), nil

}

/*
WeaviateActionsValidate validates an action based on a schema

Validate an Action's schema and meta-data. It has to be based on a schema, which is related to the given Action to be accepted by this validation.
*/
func (a *Client) WeaviateActionsValidate(params *WeaviateActionsValidateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateActionsValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateActionsValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.actions.validate",
		Method:             "POST",
		PathPattern:        "/actions/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateActionsValidateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateActionsValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
