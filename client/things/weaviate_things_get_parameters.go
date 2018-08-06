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
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateThingsGetParams creates a new WeaviateThingsGetParams object
// with the default values initialized.
func NewWeaviateThingsGetParams() *WeaviateThingsGetParams {
	var ()
	return &WeaviateThingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingsGetParamsWithTimeout creates a new WeaviateThingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingsGetParamsWithTimeout(timeout time.Duration) *WeaviateThingsGetParams {
	var ()
	return &WeaviateThingsGetParams{

		timeout: timeout,
	}
}

// NewWeaviateThingsGetParamsWithContext creates a new WeaviateThingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingsGetParamsWithContext(ctx context.Context) *WeaviateThingsGetParams {
	var ()
	return &WeaviateThingsGetParams{

		Context: ctx,
	}
}

// NewWeaviateThingsGetParamsWithHTTPClient creates a new WeaviateThingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingsGetParamsWithHTTPClient(client *http.Client) *WeaviateThingsGetParams {
	var ()
	return &WeaviateThingsGetParams{
		HTTPClient: client,
	}
}

/*WeaviateThingsGetParams contains all the parameters to send to the API endpoint
for the weaviate things get operation typically these are written to a http.Request
*/
type WeaviateThingsGetParams struct {

	/*ThingID
	  Unique ID of the thing.

	*/
	ThingID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate things get params
func (o *WeaviateThingsGetParams) WithTimeout(timeout time.Duration) *WeaviateThingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate things get params
func (o *WeaviateThingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate things get params
func (o *WeaviateThingsGetParams) WithContext(ctx context.Context) *WeaviateThingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate things get params
func (o *WeaviateThingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate things get params
func (o *WeaviateThingsGetParams) WithHTTPClient(client *http.Client) *WeaviateThingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate things get params
func (o *WeaviateThingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThingID adds the thingID to the weaviate things get params
func (o *WeaviateThingsGetParams) WithThingID(thingID strfmt.UUID) *WeaviateThingsGetParams {
	o.SetThingID(thingID)
	return o
}

// SetThingID adds the thingId to the weaviate things get params
func (o *WeaviateThingsGetParams) SetThingID(thingID strfmt.UUID) {
	o.ThingID = thingID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param thingId
	if err := r.SetPathParam("thingId", o.ThingID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
