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
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new things API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for things API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WeaviateThingHistoryGet gets a thing s history based on its uuid related to this key

Returns a particular thing history.
*/
func (a *Client) WeaviateThingHistoryGet(params *WeaviateThingHistoryGetParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingHistoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingHistoryGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.thing.history.get",
		Method:             "GET",
		PathPattern:        "/things/{thingId}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingHistoryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingHistoryGetOK), nil

}

/*
WeaviateThingsActionsList gets a thing based on its uuid related to this thing also available as websocket

Lists all actions in reverse order of creation, related to the thing that belongs to the used thingId.
*/
func (a *Client) WeaviateThingsActionsList(params *WeaviateThingsActionsListParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsActionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsActionsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.actions.list",
		Method:             "GET",
		PathPattern:        "/things/{thingId}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsActionsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsActionsListOK), nil

}

/*
WeaviateThingsCreate creates a new thing based on a thing template related to this key

Registers a new thing. Given meta-data and schema values are validated.
*/
func (a *Client) WeaviateThingsCreate(params *WeaviateThingsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsCreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.create",
		Method:             "POST",
		PathPattern:        "/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsCreateAccepted), nil

}

/*
WeaviateThingsDelete deletes a thing based on its uuid related to this key

Deletes a thing from the system. All actions pointing to this thing, where the thing is the object of the action, are also being deleted.
*/
func (a *Client) WeaviateThingsDelete(params *WeaviateThingsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.delete",
		Method:             "DELETE",
		PathPattern:        "/things/{thingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsDeleteNoContent), nil

}

/*
WeaviateThingsGet gets a thing based on its uuid related to this key

Returns a particular thing data.
*/
func (a *Client) WeaviateThingsGet(params *WeaviateThingsGetParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.get",
		Method:             "GET",
		PathPattern:        "/things/{thingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsGetOK), nil

}

/*
WeaviateThingsList gets a list of things related to this key

Lists all things in reverse order of creation, owned by the user that belongs to the used token.
*/
func (a *Client) WeaviateThingsList(params *WeaviateThingsListParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.list",
		Method:             "GET",
		PathPattern:        "/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsListOK), nil

}

/*
WeaviateThingsPatch updates a thing based on its uuid using patch semantics related to this key

Updates a thing data. This method supports patch semantics. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) WeaviateThingsPatch(params *WeaviateThingsPatchParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsPatchAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.patch",
		Method:             "PATCH",
		PathPattern:        "/things/{thingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsPatchAccepted), nil

}

/*
WeaviateThingsUpdate updates a thing based on its uuid related to this key

Updates a thing data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) WeaviateThingsUpdate(params *WeaviateThingsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsUpdateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.update",
		Method:             "PUT",
		PathPattern:        "/things/{thingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsUpdateAccepted), nil

}

/*
WeaviateThingsValidate validates things schema

Validate a thing's schema and meta-data. It has to be based on a schema, which is related to the given Thing to be accepted by this validation.
*/
func (a *Client) WeaviateThingsValidate(params *WeaviateThingsValidateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.validate",
		Method:             "POST",
		PathPattern:        "/things/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsValidateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
