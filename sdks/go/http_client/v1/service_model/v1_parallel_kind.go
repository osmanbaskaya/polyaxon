// Copyright 2018-2020 Polyaxon, Inc.
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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1ParallelKind v1 parallel kind
//
// swagger:model v1ParallelKind
type V1ParallelKind string

const (

	// V1ParallelKindRandom captures enum value "random"
	V1ParallelKindRandom V1ParallelKind = "random"

	// V1ParallelKindGrid captures enum value "grid"
	V1ParallelKindGrid V1ParallelKind = "grid"

	// V1ParallelKindHyperband captures enum value "hyperband"
	V1ParallelKindHyperband V1ParallelKind = "hyperband"

	// V1ParallelKindBayes captures enum value "bayes"
	V1ParallelKindBayes V1ParallelKind = "bayes"

	// V1ParallelKindHyperopt captures enum value "hyperopt"
	V1ParallelKindHyperopt V1ParallelKind = "hyperopt"

	// V1ParallelKindIterative captures enum value "iterative"
	V1ParallelKindIterative V1ParallelKind = "iterative"

	// V1ParallelKindMapping captures enum value "mapping"
	V1ParallelKindMapping V1ParallelKind = "mapping"
)

// for schema
var v1ParallelKindEnum []interface{}

func init() {
	var res []V1ParallelKind
	if err := json.Unmarshal([]byte(`["random","grid","hyperband","bayes","hyperopt","iterative","mapping"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ParallelKindEnum = append(v1ParallelKindEnum, v)
	}
}

func (m V1ParallelKind) validateV1ParallelKindEnum(path, location string, value V1ParallelKind) error {
	if err := validate.Enum(path, location, value, v1ParallelKindEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 parallel kind
func (m V1ParallelKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ParallelKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
