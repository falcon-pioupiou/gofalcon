// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// NewCreateAzureSubscriptionParams creates a new CreateAzureSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAzureSubscriptionParams() *CreateAzureSubscriptionParams {
	return &CreateAzureSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureSubscriptionParamsWithTimeout creates a new CreateAzureSubscriptionParams object
// with the ability to set a timeout on a request.
func NewCreateAzureSubscriptionParamsWithTimeout(timeout time.Duration) *CreateAzureSubscriptionParams {
	return &CreateAzureSubscriptionParams{
		timeout: timeout,
	}
}

// NewCreateAzureSubscriptionParamsWithContext creates a new CreateAzureSubscriptionParams object
// with the ability to set a context for a request.
func NewCreateAzureSubscriptionParamsWithContext(ctx context.Context) *CreateAzureSubscriptionParams {
	return &CreateAzureSubscriptionParams{
		Context: ctx,
	}
}

// NewCreateAzureSubscriptionParamsWithHTTPClient creates a new CreateAzureSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAzureSubscriptionParamsWithHTTPClient(client *http.Client) *CreateAzureSubscriptionParams {
	return &CreateAzureSubscriptionParams{
		HTTPClient: client,
	}
}

/*
CreateAzureSubscriptionParams contains all the parameters to send to the API endpoint

	for the create azure subscription operation.

	Typically these are written to a http.Request.
*/
type CreateAzureSubscriptionParams struct {

	// Body.
	Body *models.K8sregCreateAzureSubReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create azure subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureSubscriptionParams) WithDefaults() *CreateAzureSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create azure subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create azure subscription params
func (o *CreateAzureSubscriptionParams) WithTimeout(timeout time.Duration) *CreateAzureSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure subscription params
func (o *CreateAzureSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure subscription params
func (o *CreateAzureSubscriptionParams) WithContext(ctx context.Context) *CreateAzureSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure subscription params
func (o *CreateAzureSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure subscription params
func (o *CreateAzureSubscriptionParams) WithHTTPClient(client *http.Client) *CreateAzureSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure subscription params
func (o *CreateAzureSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create azure subscription params
func (o *CreateAzureSubscriptionParams) WithBody(body *models.K8sregCreateAzureSubReq) *CreateAzureSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create azure subscription params
func (o *CreateAzureSubscriptionParams) SetBody(body *models.K8sregCreateAzureSubReq) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
