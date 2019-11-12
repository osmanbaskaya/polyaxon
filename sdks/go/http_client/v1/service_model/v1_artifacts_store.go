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
	"github.com/go-openapi/validate"
)

// V1ArtifactsStore Artifacts store access specification
// swagger:model v1ArtifactsStore
type V1ArtifactsStore struct {

	// Optional bucket
	Bucket string `json:"bucket,omitempty"`

	// Optional time when the entityt was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional if the entity has been deleted
	Deleted bool `json:"deleted,omitempty"`

	// Optional description
	Description string `json:"description,omitempty"`

	// Optional a flag to disable an access
	Disabled bool `json:"disabled,omitempty"`

	// Optional a flag to freeze an access
	Frozen bool `json:"frozen,omitempty"`

	// Optional host path
	HostPath string `json:"host_path,omitempty"`

	// Optional the k8s secret to use
	K8sSecret string `json:"k8s_secret,omitempty"`

	// Optional mounth path
	MountPath string `json:"mount_path,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// Optional flag to set this store to read only mode
	ReadOnly bool `json:"read_only,omitempty"`

	// Optional a readme text describing this entity
	Readme string `json:"readme,omitempty"`

	// Optional Tags of this entity
	Tags []string `json:"tags"`

	// Optional type of the store
	Type string `json:"type,omitempty"`

	// Optional last time the entity was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`

	// Optional volume claim
	VolumeClaim string `json:"volume_claim,omitempty"`
}

// Validate validates this v1 artifacts store
func (m *V1ArtifactsStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ArtifactsStore) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ArtifactsStore) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ArtifactsStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ArtifactsStore) UnmarshalBinary(b []byte) error {
	var res V1ArtifactsStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
