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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsGetReader is a Reader for the WeaviateActionsGet structure.
type WeaviateActionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateActionsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewWeaviateActionsGetNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsGetOK creates a WeaviateActionsGetOK with default headers values
func NewWeaviateActionsGetOK() *WeaviateActionsGetOK {
	return &WeaviateActionsGetOK{}
}

/*WeaviateActionsGetOK handles this case with default header values.

Successful response.
*/
type WeaviateActionsGetOK struct {
	Payload *models.ActionGetResponse
}

func (o *WeaviateActionsGetOK) Error() string {
	return fmt.Sprintf("[GET /actions/{actionId}][%d] weaviateActionsGetOK  %+v", 200, o.Payload)
}

func (o *WeaviateActionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsGetUnauthorized creates a WeaviateActionsGetUnauthorized with default headers values
func NewWeaviateActionsGetUnauthorized() *WeaviateActionsGetUnauthorized {
	return &WeaviateActionsGetUnauthorized{}
}

/*WeaviateActionsGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsGetUnauthorized struct {
}

func (o *WeaviateActionsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions/{actionId}][%d] weaviateActionsGetUnauthorized ", 401)
}

func (o *WeaviateActionsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsGetForbidden creates a WeaviateActionsGetForbidden with default headers values
func NewWeaviateActionsGetForbidden() *WeaviateActionsGetForbidden {
	return &WeaviateActionsGetForbidden{}
}

/*WeaviateActionsGetForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateActionsGetForbidden struct {
}

func (o *WeaviateActionsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /actions/{actionId}][%d] weaviateActionsGetForbidden ", 403)
}

func (o *WeaviateActionsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsGetNotFound creates a WeaviateActionsGetNotFound with default headers values
func NewWeaviateActionsGetNotFound() *WeaviateActionsGetNotFound {
	return &WeaviateActionsGetNotFound{}
}

/*WeaviateActionsGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateActionsGetNotFound struct {
}

func (o *WeaviateActionsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /actions/{actionId}][%d] weaviateActionsGetNotFound ", 404)
}

func (o *WeaviateActionsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsGetNotImplemented creates a WeaviateActionsGetNotImplemented with default headers values
func NewWeaviateActionsGetNotImplemented() *WeaviateActionsGetNotImplemented {
	return &WeaviateActionsGetNotImplemented{}
}

/*WeaviateActionsGetNotImplemented handles this case with default header values.

Not (yet) implemented.
*/
type WeaviateActionsGetNotImplemented struct {
}

func (o *WeaviateActionsGetNotImplemented) Error() string {
	return fmt.Sprintf("[GET /actions/{actionId}][%d] weaviateActionsGetNotImplemented ", 501)
}

func (o *WeaviateActionsGetNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
