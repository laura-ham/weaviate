package models


// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ACLEntriesListResponse List of Access control list entries.
// swagger:model AclEntriesListResponse
type ACLEntriesListResponse struct {

	// The actual list of ACL entries.
	ACLEntries []*ACLEntry `json:"aclEntries"`

	// Identifies what kind of resource this is. Value: the fixed string "weave#aclEntriesListResponse".
	Kind *string `json:"kind,omitempty"`

	// Token corresponding to the next page of ACL entries.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// The total number of ACL entries for the query. The number of items in a response may be smaller due to paging.
	TotalResults int32 `json:"totalResults,omitempty"`
}

// Validate validates this Acl entries list response
func (m *ACLEntriesListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACLEntries(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACLEntriesListResponse) validateACLEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.ACLEntries); i++ {

		if swag.IsZero(m.ACLEntries[i]) { // not required
			continue
		}

		if m.ACLEntries[i] != nil {

			if err := m.ACLEntries[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}