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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateRunReader is a Reader for the CreateRun structure.
type CreateRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRunOK creates a CreateRunOK with default headers values
func NewCreateRunOK() *CreateRunOK {
	return &CreateRunOK{}
}

/*CreateRunOK handles this case with default header values.

A successful response.
*/
type CreateRunOK struct {
	Payload *service_model.V1Run
}

func (o *CreateRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs][%d] createRunOK  %+v", 200, o.Payload)
}

func (o *CreateRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *CreateRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunNoContent creates a CreateRunNoContent with default headers values
func NewCreateRunNoContent() *CreateRunNoContent {
	return &CreateRunNoContent{}
}

/*CreateRunNoContent handles this case with default header values.

No content.
*/
type CreateRunNoContent struct {
	Payload interface{}
}

func (o *CreateRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs][%d] createRunNoContent  %+v", 204, o.Payload)
}

func (o *CreateRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunForbidden creates a CreateRunForbidden with default headers values
func NewCreateRunForbidden() *CreateRunForbidden {
	return &CreateRunForbidden{}
}

/*CreateRunForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateRunForbidden struct {
	Payload interface{}
}

func (o *CreateRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs][%d] createRunForbidden  %+v", 403, o.Payload)
}

func (o *CreateRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunNotFound creates a CreateRunNotFound with default headers values
func NewCreateRunNotFound() *CreateRunNotFound {
	return &CreateRunNotFound{}
}

/*CreateRunNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateRunNotFound struct {
	Payload interface{}
}

func (o *CreateRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs][%d] createRunNotFound  %+v", 404, o.Payload)
}

func (o *CreateRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
