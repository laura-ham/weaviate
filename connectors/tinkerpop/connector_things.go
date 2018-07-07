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

package tinkerpop

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/creativesoftwarefdn/weaviate/connectors/utils"
	"github.com/creativesoftwarefdn/weaviate/models"
)

// AddThing adds a thing to the Foobar database with the given UUID.
// Takes the thing and a UUID as input.
// Thing is already validated against the ontology
func (f *Tinkerpop) AddThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error {

	// convert the thing to a Vertex and and Edge.
	err := f.thingOrActionToTinkerpop(UUID, thing, "thing", "add")

	// on error fail
	if err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// GetThing fills the given ThingGetResponse with the values from the database, based on the given UUID.
func (f *Tinkerpop) GetThing(ctx context.Context, UUID strfmt.UUID, thingResponse *models.ThingGetResponse) error {

	// define the ID vertex and the UUID to fetch
	result, err := f.client.Execute(
		`g.V().has("uuid", uuid).has("type", objectType)`,
		map[string]string{"uuid": UUID.String(), "objectType": "thing"},
		map[string]string{},
	)

	// in case of error, return the error
	if err != nil {
		return err
	}

	// if there is no result, send not found by sending a nil
	if result.([]interface{})[0] == nil {
		return nil
	}

	// process the thing repsonse for a single thing. thingNo = 0, because we want the first one
	err = f.processSingleThing(result, 0, thingResponse)

	// in case of error, return the error
	if err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// GetThings fills the given ThingsListResponse with the values from the database, based on the given UUIDs.
func (f *Tinkerpop) GetThings(ctx context.Context, UUIDs []strfmt.UUID, thingResponse *models.ThingsListResponse) error {
	f.messaging.DebugMessage(fmt.Sprintf("GetThings: %s", UUIDs))

	// If success return nil, otherwise return the error
	return nil
}

// ListThings fills the given ThingsListResponse with the values from the database, based on the given parameters.
func (f *Tinkerpop) ListThings(ctx context.Context, first int, offset int, keyID strfmt.UUID, wheres []*connutils.WhereQuery, thingsResponse *models.ThingsListResponse) error {

	// find the edges (if any)
	result, err := f.client.Execute(
		`g.V().has("type", objectType).range(`+strconv.Itoa((first*offset))+`, `+strconv.Itoa(first)+`)`,
		map[string]string{"objectType": "thing"},
		map[string]string{},
	)

	// Loop over the Tinkerpop results
	for thingKey := range result.([]interface{})[0].([]interface{}) {

		// define singleThing
		var singleThing models.ThingGetResponse

		// get the individual thing
		err := f.processSingleThing(result, thingKey, &singleThing)

		// in case of error, return the error
		if err != nil {
			return err
		}

		// add the thing to the response
		thingsResponse.Things = append(thingsResponse.Things, &singleThing)
	}

	// add the final results
	thingsResponse.TotalResults = int64(len(result.([]interface{})[0].([]interface{})))

	// in case of error, return the error
	if err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// UpdateThing updates the Thing in the DB at the given UUID.
func (f *Tinkerpop) UpdateThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error {

	// get the vertexes
	err := f.thingOrActionToTinkerpop(UUID, thing, "thing", "update")

	if err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// DeleteThing deletes the Thing in the DB at the given UUID.
func (f *Tinkerpop) DeleteThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error {

	// Remove based on type and uuid
	_, err := f.client.Execute(
		`g.V().has("uuid", uuid).has("type", objectType).drop()`,
		map[string]string{"uuid": UUID.String(), "objectType": "thing"},
		map[string]string{},
	)

	// return error
	if err != nil {
		return err
	}

	// If success return nil, otherwise return the error
	return nil
}

// HistoryThing fills the history of a thing based on its UUID
func (f *Tinkerpop) HistoryThing(ctx context.Context, UUID strfmt.UUID, history *models.ThingHistory) error {
	return nil
}

// MoveToHistoryThing moves a thing to history
func (f *Tinkerpop) MoveToHistoryThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID, deleted bool) error {
	return nil
}
