// Code generated by go-swagger; DO NOT EDIT.

package ods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewCancelScansParams creates a new CancelScansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelScansParams() *CancelScansParams {
	return &CancelScansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelScansParamsWithTimeout creates a new CancelScansParams object
// with the ability to set a timeout on a request.
func NewCancelScansParamsWithTimeout(timeout time.Duration) *CancelScansParams {
	return &CancelScansParams{
		timeout: timeout,
	}
}

// NewCancelScansParamsWithContext creates a new CancelScansParams object
// with the ability to set a context for a request.
func NewCancelScansParamsWithContext(ctx context.Context) *CancelScansParams {
	return &CancelScansParams{
		Context: ctx,
	}
}

// NewCancelScansParamsWithHTTPClient creates a new CancelScansParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelScansParamsWithHTTPClient(client *http.Client) *CancelScansParams {
	return &CancelScansParams{
		HTTPClient: client,
	}
}

/*
CancelScansParams contains all the parameters to send to the API endpoint

	for the cancel scans operation.

	Typically these are written to a http.Request.
*/
type CancelScansParams struct {

	// Body.
	Body *models.EntitiesODSCancelScanRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel scans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelScansParams) WithDefaults() *CancelScansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel scans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelScansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel scans params
func (o *CancelScansParams) WithTimeout(timeout time.Duration) *CancelScansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel scans params
func (o *CancelScansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel scans params
func (o *CancelScansParams) WithContext(ctx context.Context) *CancelScansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel scans params
func (o *CancelScansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel scans params
func (o *CancelScansParams) WithHTTPClient(client *http.Client) *CancelScansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel scans params
func (o *CancelScansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cancel scans params
func (o *CancelScansParams) WithBody(body *models.EntitiesODSCancelScanRequest) *CancelScansParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cancel scans params
func (o *CancelScansParams) SetBody(body *models.EntitiesODSCancelScanRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CancelScansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}