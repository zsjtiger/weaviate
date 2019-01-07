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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateActionsPropertiesDeleteHandlerFunc turns a function with the right signature into a weaviate actions properties delete handler
type WeaviateActionsPropertiesDeleteHandlerFunc func(WeaviateActionsPropertiesDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsPropertiesDeleteHandlerFunc) Handle(params WeaviateActionsPropertiesDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsPropertiesDeleteHandler interface for that can handle valid weaviate actions properties delete params
type WeaviateActionsPropertiesDeleteHandler interface {
	Handle(WeaviateActionsPropertiesDeleteParams, interface{}) middleware.Responder
}

// NewWeaviateActionsPropertiesDelete creates a new http.Handler for the weaviate actions properties delete operation
func NewWeaviateActionsPropertiesDelete(ctx *middleware.Context, handler WeaviateActionsPropertiesDeleteHandler) *WeaviateActionsPropertiesDelete {
	return &WeaviateActionsPropertiesDelete{Context: ctx, Handler: handler}
}

/*WeaviateActionsPropertiesDelete swagger:route DELETE /actions/{actionId}/properties/{propertyName} actions weaviateActionsPropertiesDelete

Delete the single reference that is given in the body from the list of references that this property has.

Delete the single reference that is given in the body from the list of references that this property has.

*/
type WeaviateActionsPropertiesDelete struct {
	Context *middleware.Context
	Handler WeaviateActionsPropertiesDeleteHandler
}

func (o *WeaviateActionsPropertiesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsPropertiesDeleteParams()

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
