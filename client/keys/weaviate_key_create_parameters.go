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

package keys

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

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateKeyCreateParams creates a new WeaviateKeyCreateParams object
// with the default values initialized.
func NewWeaviateKeyCreateParams() *WeaviateKeyCreateParams {
	var ()
	return &WeaviateKeyCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateKeyCreateParamsWithTimeout creates a new WeaviateKeyCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateKeyCreateParamsWithTimeout(timeout time.Duration) *WeaviateKeyCreateParams {
	var ()
	return &WeaviateKeyCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateKeyCreateParamsWithContext creates a new WeaviateKeyCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateKeyCreateParamsWithContext(ctx context.Context) *WeaviateKeyCreateParams {
	var ()
	return &WeaviateKeyCreateParams{

		Context: ctx,
	}
}

// NewWeaviateKeyCreateParamsWithHTTPClient creates a new WeaviateKeyCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateKeyCreateParamsWithHTTPClient(client *http.Client) *WeaviateKeyCreateParams {
	var ()
	return &WeaviateKeyCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateKeyCreateParams contains all the parameters to send to the API endpoint
for the weaviate key create operation typically these are written to a http.Request
*/
type WeaviateKeyCreateParams struct {

	/*Body*/
	Body *models.KeyCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate key create params
func (o *WeaviateKeyCreateParams) WithTimeout(timeout time.Duration) *WeaviateKeyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate key create params
func (o *WeaviateKeyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate key create params
func (o *WeaviateKeyCreateParams) WithContext(ctx context.Context) *WeaviateKeyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate key create params
func (o *WeaviateKeyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate key create params
func (o *WeaviateKeyCreateParams) WithHTTPClient(client *http.Client) *WeaviateKeyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate key create params
func (o *WeaviateKeyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate key create params
func (o *WeaviateKeyCreateParams) WithBody(body *models.KeyCreate) *WeaviateKeyCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate key create params
func (o *WeaviateKeyCreateParams) SetBody(body *models.KeyCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateKeyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
