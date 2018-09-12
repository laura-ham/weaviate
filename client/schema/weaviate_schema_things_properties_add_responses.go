// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaThingsPropertiesAddReader is a Reader for the WeaviateSchemaThingsPropertiesAdd structure.
type WeaviateSchemaThingsPropertiesAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsPropertiesAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsPropertiesAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaThingsPropertiesAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaThingsPropertiesAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaThingsPropertiesAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsPropertiesAddOK creates a WeaviateSchemaThingsPropertiesAddOK with default headers values
func NewWeaviateSchemaThingsPropertiesAddOK() *WeaviateSchemaThingsPropertiesAddOK {
	return &WeaviateSchemaThingsPropertiesAddOK{}
}

/*WeaviateSchemaThingsPropertiesAddOK handles this case with default header values.

Added the property
*/
type WeaviateSchemaThingsPropertiesAddOK struct {
}

func (o *WeaviateSchemaThingsPropertiesAddOK) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] weaviateSchemaThingsPropertiesAddOK ", 200)
}

func (o *WeaviateSchemaThingsPropertiesAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesAddUnauthorized creates a WeaviateSchemaThingsPropertiesAddUnauthorized with default headers values
func NewWeaviateSchemaThingsPropertiesAddUnauthorized() *WeaviateSchemaThingsPropertiesAddUnauthorized {
	return &WeaviateSchemaThingsPropertiesAddUnauthorized{}
}

/*WeaviateSchemaThingsPropertiesAddUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsPropertiesAddUnauthorized struct {
}

func (o *WeaviateSchemaThingsPropertiesAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] weaviateSchemaThingsPropertiesAddUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsPropertiesAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesAddForbidden creates a WeaviateSchemaThingsPropertiesAddForbidden with default headers values
func NewWeaviateSchemaThingsPropertiesAddForbidden() *WeaviateSchemaThingsPropertiesAddForbidden {
	return &WeaviateSchemaThingsPropertiesAddForbidden{}
}

/*WeaviateSchemaThingsPropertiesAddForbidden handles this case with default header values.

Could not find the Thing class
*/
type WeaviateSchemaThingsPropertiesAddForbidden struct {
}

func (o *WeaviateSchemaThingsPropertiesAddForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] weaviateSchemaThingsPropertiesAddForbidden ", 403)
}

func (o *WeaviateSchemaThingsPropertiesAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesAddUnprocessableEntity creates a WeaviateSchemaThingsPropertiesAddUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsPropertiesAddUnprocessableEntity() *WeaviateSchemaThingsPropertiesAddUnprocessableEntity {
	return &WeaviateSchemaThingsPropertiesAddUnprocessableEntity{}
}

/*WeaviateSchemaThingsPropertiesAddUnprocessableEntity handles this case with default header values.

Invalid property
*/
type WeaviateSchemaThingsPropertiesAddUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsPropertiesAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] weaviateSchemaThingsPropertiesAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaThingsPropertiesAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
