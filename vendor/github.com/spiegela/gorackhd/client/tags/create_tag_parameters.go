package tags

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

	"github.com/spiegela/gorackhd/models"
)

// NewCreateTagParams creates a new CreateTagParams object
// with the default values initialized.
func NewCreateTagParams() *CreateTagParams {
	var ()
	return &CreateTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTagParamsWithTimeout creates a new CreateTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTagParamsWithTimeout(timeout time.Duration) *CreateTagParams {
	var ()
	return &CreateTagParams{

		timeout: timeout,
	}
}

// NewCreateTagParamsWithContext creates a new CreateTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTagParamsWithContext(ctx context.Context) *CreateTagParams {
	var ()
	return &CreateTagParams{

		Context: ctx,
	}
}

// NewCreateTagParamsWithHTTPClient creates a new CreateTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTagParamsWithHTTPClient(client *http.Client) *CreateTagParams {
	var ()
	return &CreateTagParams{
		HTTPClient: client,
	}
}

/*CreateTagParams contains all the parameters to send to the API endpoint
for the create tag operation typically these are written to a http.Request
*/
type CreateTagParams struct {

	/*Body
	  The tag properties

	*/
	Body *models.PostTags

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create tag params
func (o *CreateTagParams) WithTimeout(timeout time.Duration) *CreateTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tag params
func (o *CreateTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tag params
func (o *CreateTagParams) WithContext(ctx context.Context) *CreateTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tag params
func (o *CreateTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tag params
func (o *CreateTagParams) WithHTTPClient(client *http.Client) *CreateTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tag params
func (o *CreateTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create tag params
func (o *CreateTagParams) WithBody(body *models.PostTags) *CreateTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create tag params
func (o *CreateTagParams) SetBody(body *models.PostTags) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.PostTags)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}