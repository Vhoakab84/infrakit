package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// SkusIDPutPackReader is a Reader for the SkusIDPutPack structure.
type SkusIDPutPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SkusIDPutPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSkusIDPutPackCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSkusIDPutPackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSkusIDPutPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSkusIDPutPackCreated creates a SkusIDPutPackCreated with default headers values
func NewSkusIDPutPackCreated() *SkusIDPutPackCreated {
	return &SkusIDPutPackCreated{}
}

/*SkusIDPutPackCreated handles this case with default header values.

Successfully created the SKU Pack
*/
type SkusIDPutPackCreated struct {
	Payload SkusIDPutPackCreatedBody
}

func (o *SkusIDPutPackCreated) Error() string {
	return fmt.Sprintf("[PUT /skus/{identifier}/pack][%d] skusIdPutPackCreated  %+v", 201, o.Payload)
}

func (o *SkusIDPutPackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusIDPutPackNotFound creates a SkusIDPutPackNotFound with default headers values
func NewSkusIDPutPackNotFound() *SkusIDPutPackNotFound {
	return &SkusIDPutPackNotFound{}
}

/*SkusIDPutPackNotFound handles this case with default header values.

The SKU with the specified identifier was not found
*/
type SkusIDPutPackNotFound struct {
	Payload *models.Error
}

func (o *SkusIDPutPackNotFound) Error() string {
	return fmt.Sprintf("[PUT /skus/{identifier}/pack][%d] skusIdPutPackNotFound  %+v", 404, o.Payload)
}

func (o *SkusIDPutPackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusIDPutPackDefault creates a SkusIDPutPackDefault with default headers values
func NewSkusIDPutPackDefault(code int) *SkusIDPutPackDefault {
	return &SkusIDPutPackDefault{
		_statusCode: code,
	}
}

/*SkusIDPutPackDefault handles this case with default header values.

Unexpected error
*/
type SkusIDPutPackDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the skus Id put pack default response
func (o *SkusIDPutPackDefault) Code() int {
	return o._statusCode
}

func (o *SkusIDPutPackDefault) Error() string {
	return fmt.Sprintf("[PUT /skus/{identifier}/pack][%d] skusIdPutPack default  %+v", o._statusCode, o.Payload)
}

func (o *SkusIDPutPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SkusIDPutPackCreatedBody skus ID put pack created body
swagger:model SkusIDPutPackCreatedBody
*/
type SkusIDPutPackCreatedBody interface{}