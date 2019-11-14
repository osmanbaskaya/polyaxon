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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRestoreProjectParams creates a new RestoreProjectParams object
// with the default values initialized.
func NewRestoreProjectParams() *RestoreProjectParams {
	var ()
	return &RestoreProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreProjectParamsWithTimeout creates a new RestoreProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestoreProjectParamsWithTimeout(timeout time.Duration) *RestoreProjectParams {
	var ()
	return &RestoreProjectParams{

		timeout: timeout,
	}
}

// NewRestoreProjectParamsWithContext creates a new RestoreProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestoreProjectParamsWithContext(ctx context.Context) *RestoreProjectParams {
	var ()
	return &RestoreProjectParams{

		Context: ctx,
	}
}

// NewRestoreProjectParamsWithHTTPClient creates a new RestoreProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestoreProjectParamsWithHTTPClient(client *http.Client) *RestoreProjectParams {
	var ()
	return &RestoreProjectParams{
		HTTPClient: client,
	}
}

/*RestoreProjectParams contains all the parameters to send to the API endpoint
for the restore project operation typically these are written to a http.Request
*/
type RestoreProjectParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restore project params
func (o *RestoreProjectParams) WithTimeout(timeout time.Duration) *RestoreProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore project params
func (o *RestoreProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore project params
func (o *RestoreProjectParams) WithContext(ctx context.Context) *RestoreProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore project params
func (o *RestoreProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore project params
func (o *RestoreProjectParams) WithHTTPClient(client *http.Client) *RestoreProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore project params
func (o *RestoreProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the restore project params
func (o *RestoreProjectParams) WithOwner(owner string) *RestoreProjectParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the restore project params
func (o *RestoreProjectParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the restore project params
func (o *RestoreProjectParams) WithProject(project string) *RestoreProjectParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the restore project params
func (o *RestoreProjectParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}