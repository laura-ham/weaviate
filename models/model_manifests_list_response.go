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
 package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ModelManifestsListResponse List of model manifests.
// swagger:model ModelManifestsListResponse
type ModelManifestsListResponse struct {

	// Identifies what kind of resource this is. Value: the fixed string "weave#modelManifestsListResponse".
	Kind *string `json:"kind,omitempty"`

	// The actual list of model manifests.
	ModelManifests []*ModelManifest `json:"modelManifests"`

	// Token corresponding to the next page of model manifests.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// The total number of model manifests for the query. The number of items in a response may be smaller due to paging.
	TotalResults int32 `json:"totalResults,omitempty"`
}

// Validate validates this model manifests list response
func (m *ModelManifestsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModelManifests(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelManifestsListResponse) validateModelManifests(formats strfmt.Registry) error {

	if swag.IsZero(m.ModelManifests) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelManifests); i++ {

		if swag.IsZero(m.ModelManifests[i]) { // not required
			continue
		}

		if m.ModelManifests[i] != nil {

			if err := m.ModelManifests[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}