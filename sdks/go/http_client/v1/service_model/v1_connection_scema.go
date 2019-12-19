// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1ConnectionScema v1 connection scema
// swagger:model v1ConnectionScema
type V1ConnectionScema struct {

	// blob connection
	BlobConnection *V1BlobConnection `json:"blob_connection,omitempty"`

	// claim connection
	ClaimConnection *V1ClaimConnection `json:"claim_connection,omitempty"`

	// host connection
	HostConnection *V1HostConnection `json:"host_connection,omitempty"`

	// host path connection
	HostPathConnection *V1HostPathConnection `json:"host_path_connection,omitempty"`
}

// Validate validates this v1 connection scema
func (m *V1ConnectionScema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlobConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPathConnection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ConnectionScema) validateBlobConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.BlobConnection) { // not required
		return nil
	}

	if m.BlobConnection != nil {
		if err := m.BlobConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blob_connection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionScema) validateClaimConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.ClaimConnection) { // not required
		return nil
	}

	if m.ClaimConnection != nil {
		if err := m.ClaimConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claim_connection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionScema) validateHostConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.HostConnection) { // not required
		return nil
	}

	if m.HostConnection != nil {
		if err := m.HostConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host_connection")
			}
			return err
		}
	}

	return nil
}

func (m *V1ConnectionScema) validateHostPathConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.HostPathConnection) { // not required
		return nil
	}

	if m.HostPathConnection != nil {
		if err := m.HostPathConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host_path_connection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ConnectionScema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ConnectionScema) UnmarshalBinary(b []byte) error {
	var res V1ConnectionScema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}