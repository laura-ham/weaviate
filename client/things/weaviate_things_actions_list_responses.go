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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsActionsListReader is a Reader for the WeaviateThingsActionsList structure.
type WeaviateThingsActionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsActionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsActionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsActionsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsActionsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsActionsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewWeaviateThingsActionsListNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsActionsListOK creates a WeaviateThingsActionsListOK with default headers values
func NewWeaviateThingsActionsListOK() *WeaviateThingsActionsListOK {
	return &WeaviateThingsActionsListOK{}
}

/*WeaviateThingsActionsListOK handles this case with default header values.

Successful response.
*/
type WeaviateThingsActionsListOK struct {
	Payload *models.ActionsListResponse
}

func (o *WeaviateThingsActionsListOK) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/actions][%d] weaviateThingsActionsListOK  %+v", 200, o.Payload)
}

func (o *WeaviateThingsActionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsActionsListUnauthorized creates a WeaviateThingsActionsListUnauthorized with default headers values
func NewWeaviateThingsActionsListUnauthorized() *WeaviateThingsActionsListUnauthorized {
	return &WeaviateThingsActionsListUnauthorized{}
}

/*WeaviateThingsActionsListUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsActionsListUnauthorized struct {
}

func (o *WeaviateThingsActionsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/actions][%d] weaviateThingsActionsListUnauthorized ", 401)
}

func (o *WeaviateThingsActionsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsActionsListForbidden creates a WeaviateThingsActionsListForbidden with default headers values
func NewWeaviateThingsActionsListForbidden() *WeaviateThingsActionsListForbidden {
	return &WeaviateThingsActionsListForbidden{}
}

/*WeaviateThingsActionsListForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsActionsListForbidden struct {
}

func (o *WeaviateThingsActionsListForbidden) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/actions][%d] weaviateThingsActionsListForbidden ", 403)
}

func (o *WeaviateThingsActionsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsActionsListNotFound creates a WeaviateThingsActionsListNotFound with default headers values
func NewWeaviateThingsActionsListNotFound() *WeaviateThingsActionsListNotFound {
	return &WeaviateThingsActionsListNotFound{}
}

/*WeaviateThingsActionsListNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsActionsListNotFound struct {
}

func (o *WeaviateThingsActionsListNotFound) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/actions][%d] weaviateThingsActionsListNotFound ", 404)
}

func (o *WeaviateThingsActionsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsActionsListNotImplemented creates a WeaviateThingsActionsListNotImplemented with default headers values
func NewWeaviateThingsActionsListNotImplemented() *WeaviateThingsActionsListNotImplemented {
	return &WeaviateThingsActionsListNotImplemented{}
}

/*WeaviateThingsActionsListNotImplemented handles this case with default header values.

Not (yet) implemented.
*/
type WeaviateThingsActionsListNotImplemented struct {
}

func (o *WeaviateThingsActionsListNotImplemented) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/actions][%d] weaviateThingsActionsListNotImplemented ", 501)
}

func (o *WeaviateThingsActionsListNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
