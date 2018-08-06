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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/genesis/models"
)

// GenesisPeersListReader is a Reader for the GenesisPeersList structure.
type GenesisPeersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenesisPeersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenesisPeersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGenesisPeersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenesisPeersListOK creates a GenesisPeersListOK with default headers values
func NewGenesisPeersListOK() *GenesisPeersListOK {
	return &GenesisPeersListOK{}
}

/*GenesisPeersListOK handles this case with default header values.

The list of registered peers
*/
type GenesisPeersListOK struct {
	Payload []*models.Peer
}

func (o *GenesisPeersListOK) Error() string {
	return fmt.Sprintf("[GET /peers][%d] genesisPeersListOK  %+v", 200, o.Payload)
}

func (o *GenesisPeersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenesisPeersListInternalServerError creates a GenesisPeersListInternalServerError with default headers values
func NewGenesisPeersListInternalServerError() *GenesisPeersListInternalServerError {
	return &GenesisPeersListInternalServerError{}
}

/*GenesisPeersListInternalServerError handles this case with default header values.

Internal error
*/
type GenesisPeersListInternalServerError struct {
}

func (o *GenesisPeersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /peers][%d] genesisPeersListInternalServerError ", 500)
}

func (o *GenesisPeersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
