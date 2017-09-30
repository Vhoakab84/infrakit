package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ProfilesGetReader is a Reader for the ProfilesGet structure.
type ProfilesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProfilesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewProfilesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProfilesGetOK creates a ProfilesGetOK with default headers values
func NewProfilesGetOK() *ProfilesGetOK {
	return &ProfilesGetOK{}
}

/*ProfilesGetOK handles this case with default header values.

Successfully retrieved a list of profiles for specified mac / ip
*/
type ProfilesGetOK struct {
	Payload ProfilesGetOKBody
}

func (o *ProfilesGetOK) Error() string {
	return fmt.Sprintf("[GET /profiles][%d] profilesGetOK  %+v", 200, o.Payload)
}

func (o *ProfilesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesGetDefault creates a ProfilesGetDefault with default headers values
func NewProfilesGetDefault(code int) *ProfilesGetDefault {
	return &ProfilesGetDefault{
		_statusCode: code,
	}
}

/*ProfilesGetDefault handles this case with default header values.

Unexpected error
*/
type ProfilesGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the profiles get default response
func (o *ProfilesGetDefault) Code() int {
	return o._statusCode
}

func (o *ProfilesGetDefault) Error() string {
	return fmt.Sprintf("[GET /profiles][%d] profilesGet default  %+v", o._statusCode, o.Payload)
}

func (o *ProfilesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProfilesGetOKBody profiles get o k body
swagger:model ProfilesGetOKBody
*/
type ProfilesGetOKBody interface{}