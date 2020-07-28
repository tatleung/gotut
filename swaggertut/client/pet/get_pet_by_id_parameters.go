// Code generated by go-swagger; DO NOT EDIT.

package pet

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
	"github.com/go-openapi/swag"
)

// NewGetPetByIDParams creates a new GetPetByIDParams object
// with the default values initialized.
func NewGetPetByIDParams() *GetPetByIDParams {
	var ()
	return &GetPetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPetByIDParamsWithTimeout creates a new GetPetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPetByIDParamsWithTimeout(timeout time.Duration) *GetPetByIDParams {
	var ()
	return &GetPetByIDParams{

		timeout: timeout,
	}
}

// NewGetPetByIDParamsWithContext creates a new GetPetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPetByIDParamsWithContext(ctx context.Context) *GetPetByIDParams {
	var ()
	return &GetPetByIDParams{

		Context: ctx,
	}
}

// NewGetPetByIDParamsWithHTTPClient creates a new GetPetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPetByIDParamsWithHTTPClient(client *http.Client) *GetPetByIDParams {
	var ()
	return &GetPetByIDParams{
		HTTPClient: client,
	}
}

/*GetPetByIDParams contains all the parameters to send to the API endpoint
for the get pet by Id operation typically these are written to a http.Request
*/
type GetPetByIDParams struct {

	/*PetID
	  ID of pet to return

	*/
	PetID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pet by Id params
func (o *GetPetByIDParams) WithTimeout(timeout time.Duration) *GetPetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pet by Id params
func (o *GetPetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pet by Id params
func (o *GetPetByIDParams) WithContext(ctx context.Context) *GetPetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pet by Id params
func (o *GetPetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pet by Id params
func (o *GetPetByIDParams) WithHTTPClient(client *http.Client) *GetPetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pet by Id params
func (o *GetPetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPetID adds the petID to the get pet by Id params
func (o *GetPetByIDParams) WithPetID(petID int64) *GetPetByIDParams {
	o.SetPetID(petID)
	return o
}

// SetPetID adds the petId to the get pet by Id params
func (o *GetPetByIDParams) SetPetID(petID int64) {
	o.PetID = petID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param petId
	if err := r.SetPathParam("petId", swag.FormatInt64(o.PetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
