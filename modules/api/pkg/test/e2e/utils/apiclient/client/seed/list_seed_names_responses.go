// Code generated by go-swagger; DO NOT EDIT.

package seed

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListSeedNamesReader is a Reader for the ListSeedNames structure.
type ListSeedNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSeedNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSeedNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSeedNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSeedNamesOK creates a ListSeedNamesOK with default headers values
func NewListSeedNamesOK() *ListSeedNamesOK {
	return &ListSeedNamesOK{}
}

/*
ListSeedNamesOK describes a response with status code 200, with default header values.

SeedNamesList
*/
type ListSeedNamesOK struct {
	Payload models.SeedNamesList
}

// IsSuccess returns true when this list seed names o k response has a 2xx status code
func (o *ListSeedNamesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list seed names o k response has a 3xx status code
func (o *ListSeedNamesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list seed names o k response has a 4xx status code
func (o *ListSeedNamesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list seed names o k response has a 5xx status code
func (o *ListSeedNamesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list seed names o k response a status code equal to that given
func (o *ListSeedNamesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListSeedNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/seed][%d] listSeedNamesOK  %+v", 200, o.Payload)
}

func (o *ListSeedNamesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/seed][%d] listSeedNamesOK  %+v", 200, o.Payload)
}

func (o *ListSeedNamesOK) GetPayload() models.SeedNamesList {
	return o.Payload
}

func (o *ListSeedNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSeedNamesDefault creates a ListSeedNamesDefault with default headers values
func NewListSeedNamesDefault(code int) *ListSeedNamesDefault {
	return &ListSeedNamesDefault{
		_statusCode: code,
	}
}

/*
ListSeedNamesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListSeedNamesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list seed names default response
func (o *ListSeedNamesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list seed names default response has a 2xx status code
func (o *ListSeedNamesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list seed names default response has a 3xx status code
func (o *ListSeedNamesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list seed names default response has a 4xx status code
func (o *ListSeedNamesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list seed names default response has a 5xx status code
func (o *ListSeedNamesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list seed names default response a status code equal to that given
func (o *ListSeedNamesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListSeedNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/seed][%d] listSeedNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListSeedNamesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/seed][%d] listSeedNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListSeedNamesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSeedNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}