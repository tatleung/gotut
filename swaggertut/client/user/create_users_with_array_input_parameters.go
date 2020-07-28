// Code generated by go-swagger; DO NOT EDIT.

package user

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

	"github.com/tatleung/gotut/swaggertut/models"
)

// NewCreateUsersWithArrayInputParams creates a new CreateUsersWithArrayInputParams object
// with the default values initialized.
func NewCreateUsersWithArrayInputParams() *CreateUsersWithArrayInputParams {
	var ()
	return &CreateUsersWithArrayInputParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsersWithArrayInputParamsWithTimeout creates a new CreateUsersWithArrayInputParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUsersWithArrayInputParamsWithTimeout(timeout time.Duration) *CreateUsersWithArrayInputParams {
	var ()
	return &CreateUsersWithArrayInputParams{

		timeout: timeout,
	}
}

// NewCreateUsersWithArrayInputParamsWithContext creates a new CreateUsersWithArrayInputParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUsersWithArrayInputParamsWithContext(ctx context.Context) *CreateUsersWithArrayInputParams {
	var ()
	return &CreateUsersWithArrayInputParams{

		Context: ctx,
	}
}

// NewCreateUsersWithArrayInputParamsWithHTTPClient creates a new CreateUsersWithArrayInputParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUsersWithArrayInputParamsWithHTTPClient(client *http.Client) *CreateUsersWithArrayInputParams {
	var ()
	return &CreateUsersWithArrayInputParams{
		HTTPClient: client,
	}
}

/*CreateUsersWithArrayInputParams contains all the parameters to send to the API endpoint
for the create users with array input operation typically these are written to a http.Request
*/
type CreateUsersWithArrayInputParams struct {

	/*Body
	  List of user object

	*/
	Body []*models.User

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create users with array input params
func (o *CreateUsersWithArrayInputParams) WithTimeout(timeout time.Duration) *CreateUsersWithArrayInputParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create users with array input params
func (o *CreateUsersWithArrayInputParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create users with array input params
func (o *CreateUsersWithArrayInputParams) WithContext(ctx context.Context) *CreateUsersWithArrayInputParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create users with array input params
func (o *CreateUsersWithArrayInputParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create users with array input params
func (o *CreateUsersWithArrayInputParams) WithHTTPClient(client *http.Client) *CreateUsersWithArrayInputParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create users with array input params
func (o *CreateUsersWithArrayInputParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create users with array input params
func (o *CreateUsersWithArrayInputParams) WithBody(body []*models.User) *CreateUsersWithArrayInputParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create users with array input params
func (o *CreateUsersWithArrayInputParams) SetBody(body []*models.User) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsersWithArrayInputParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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