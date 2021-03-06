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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateThingsDeleteNoContentCode is the HTTP code returned for type WeaviateThingsDeleteNoContent
const WeaviateThingsDeleteNoContentCode int = 204

/*WeaviateThingsDeleteNoContent Successful deleted.

swagger:response weaviateThingsDeleteNoContent
*/
type WeaviateThingsDeleteNoContent struct {
}

// NewWeaviateThingsDeleteNoContent creates WeaviateThingsDeleteNoContent with default headers values
func NewWeaviateThingsDeleteNoContent() *WeaviateThingsDeleteNoContent {

	return &WeaviateThingsDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateThingsDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// WeaviateThingsDeleteUnauthorizedCode is the HTTP code returned for type WeaviateThingsDeleteUnauthorized
const WeaviateThingsDeleteUnauthorizedCode int = 401

/*WeaviateThingsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsDeleteUnauthorized
*/
type WeaviateThingsDeleteUnauthorized struct {
}

// NewWeaviateThingsDeleteUnauthorized creates WeaviateThingsDeleteUnauthorized with default headers values
func NewWeaviateThingsDeleteUnauthorized() *WeaviateThingsDeleteUnauthorized {

	return &WeaviateThingsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateThingsDeleteForbiddenCode is the HTTP code returned for type WeaviateThingsDeleteForbidden
const WeaviateThingsDeleteForbiddenCode int = 403

/*WeaviateThingsDeleteForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsDeleteForbidden
*/
type WeaviateThingsDeleteForbidden struct {
}

// NewWeaviateThingsDeleteForbidden creates WeaviateThingsDeleteForbidden with default headers values
func NewWeaviateThingsDeleteForbidden() *WeaviateThingsDeleteForbidden {

	return &WeaviateThingsDeleteForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateThingsDeleteNotFoundCode is the HTTP code returned for type WeaviateThingsDeleteNotFound
const WeaviateThingsDeleteNotFoundCode int = 404

/*WeaviateThingsDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateThingsDeleteNotFound
*/
type WeaviateThingsDeleteNotFound struct {
}

// NewWeaviateThingsDeleteNotFound creates WeaviateThingsDeleteNotFound with default headers values
func NewWeaviateThingsDeleteNotFound() *WeaviateThingsDeleteNotFound {

	return &WeaviateThingsDeleteNotFound{}
}

// WriteResponse to the client
func (o *WeaviateThingsDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
