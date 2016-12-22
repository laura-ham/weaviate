/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package devices


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveDevicesDeleteHandlerFunc turns a function with the right signature into a weave devices delete handler
type WeaveDevicesDeleteHandlerFunc func(WeaveDevicesDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveDevicesDeleteHandlerFunc) Handle(params WeaveDevicesDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveDevicesDeleteHandler interface for that can handle valid weave devices delete params
type WeaveDevicesDeleteHandler interface {
	Handle(WeaveDevicesDeleteParams, interface{}) middleware.Responder
}

// NewWeaveDevicesDelete creates a new http.Handler for the weave devices delete operation
func NewWeaveDevicesDelete(ctx *middleware.Context, handler WeaveDevicesDeleteHandler) *WeaveDevicesDelete {
	return &WeaveDevicesDelete{Context: ctx, Handler: handler}
}

/*WeaveDevicesDelete swagger:route DELETE /devices/{deviceId} devices weaveDevicesDelete

Deletes a device from the system.

*/
type WeaveDevicesDelete struct {
	Context *middleware.Context
	Handler WeaveDevicesDeleteHandler
}

func (o *WeaveDevicesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveDevicesDeleteParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
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