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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActionGetHistoryResponse action get history response
// swagger:model ActionGetHistoryResponse
type ActionGetHistoryResponse struct {
	ActionHistory

	// action Id
	// Format: uuid
	ActionID strfmt.UUID `json:"actionId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ActionGetHistoryResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ActionHistory
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ActionHistory = aO0

	// AO1
	var dataAO1 struct {
		ActionID strfmt.UUID `json:"actionId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ActionID = dataAO1.ActionID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ActionGetHistoryResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ActionHistory)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ActionID strfmt.UUID `json:"actionId,omitempty"`
	}

	dataAO1.ActionID = m.ActionID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this action get history response
func (m *ActionGetHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ActionHistory
	if err := m.ActionHistory.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionGetHistoryResponse) validateActionID(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionID) { // not required
		return nil
	}

	if err := validate.FormatOf("actionId", "body", "uuid", m.ActionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionGetHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionGetHistoryResponse) UnmarshalBinary(b []byte) error {
	var res ActionGetHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
