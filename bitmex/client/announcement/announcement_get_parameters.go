// Code generated by go-swagger; DO NOT EDIT.

package announcement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAnnouncementGetParams creates a new AnnouncementGetParams object
// with the default values initialized.
func NewAnnouncementGetParams() *AnnouncementGetParams {
	var ()
	return &AnnouncementGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAnnouncementGetParamsWithTimeout creates a new AnnouncementGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAnnouncementGetParamsWithTimeout(timeout time.Duration) *AnnouncementGetParams {
	var ()
	return &AnnouncementGetParams{

		timeout: timeout,
	}
}

// NewAnnouncementGetParamsWithContext creates a new AnnouncementGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAnnouncementGetParamsWithContext(ctx context.Context) *AnnouncementGetParams {
	var ()
	return &AnnouncementGetParams{

		Context: ctx,
	}
}

// NewAnnouncementGetParamsWithHTTPClient creates a new AnnouncementGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAnnouncementGetParamsWithHTTPClient(client *http.Client) *AnnouncementGetParams {
	var ()
	return &AnnouncementGetParams{
		HTTPClient: client,
	}
}

/*AnnouncementGetParams contains all the parameters to send to the API endpoint
for the announcement get operation typically these are written to a http.Request
*/
type AnnouncementGetParams struct {

	/*Columns
	  Array of column names to fetch. If omitted, will return all columns.

	*/
	Columns *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the announcement get params
func (o *AnnouncementGetParams) WithTimeout(timeout time.Duration) *AnnouncementGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the announcement get params
func (o *AnnouncementGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the announcement get params
func (o *AnnouncementGetParams) WithContext(ctx context.Context) *AnnouncementGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the announcement get params
func (o *AnnouncementGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the announcement get params
func (o *AnnouncementGetParams) WithHTTPClient(client *http.Client) *AnnouncementGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the announcement get params
func (o *AnnouncementGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumns adds the columns to the announcement get params
func (o *AnnouncementGetParams) WithColumns(columns *string) *AnnouncementGetParams {
	o.SetColumns(columns)
	return o
}

// SetColumns adds the columns to the announcement get params
func (o *AnnouncementGetParams) SetColumns(columns *string) {
	o.Columns = columns
}

// WriteToRequest writes these params to a swagger request
func (o *AnnouncementGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Columns != nil {

		// query param columns
		var qrColumns string
		if o.Columns != nil {
			qrColumns = *o.Columns
		}
		qColumns := qrColumns
		if qColumns != "" {
			if err := r.SetQueryParam("columns", qColumns); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}