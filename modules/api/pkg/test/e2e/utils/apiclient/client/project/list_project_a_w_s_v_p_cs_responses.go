// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectAWSVPCsReader is a Reader for the ListProjectAWSVPCs structure.
type ListProjectAWSVPCsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectAWSVPCsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectAWSVPCsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectAWSVPCsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectAWSVPCsOK creates a ListProjectAWSVPCsOK with default headers values
func NewListProjectAWSVPCsOK() *ListProjectAWSVPCsOK {
	return &ListProjectAWSVPCsOK{}
}

/*
ListProjectAWSVPCsOK describes a response with status code 200, with default header values.

AWSVPCList
*/
type ListProjectAWSVPCsOK struct {
	Payload models.AWSVPCList
}

// IsSuccess returns true when this list project a w s v p cs o k response has a 2xx status code
func (o *ListProjectAWSVPCsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project a w s v p cs o k response has a 3xx status code
func (o *ListProjectAWSVPCsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project a w s v p cs o k response has a 4xx status code
func (o *ListProjectAWSVPCsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project a w s v p cs o k response has a 5xx status code
func (o *ListProjectAWSVPCsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project a w s v p cs o k response a status code equal to that given
func (o *ListProjectAWSVPCsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectAWSVPCsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/vpcs][%d] listProjectAWSVPCsOK  %+v", 200, o.Payload)
}

func (o *ListProjectAWSVPCsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/vpcs][%d] listProjectAWSVPCsOK  %+v", 200, o.Payload)
}

func (o *ListProjectAWSVPCsOK) GetPayload() models.AWSVPCList {
	return o.Payload
}

func (o *ListProjectAWSVPCsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectAWSVPCsDefault creates a ListProjectAWSVPCsDefault with default headers values
func NewListProjectAWSVPCsDefault(code int) *ListProjectAWSVPCsDefault {
	return &ListProjectAWSVPCsDefault{
		_statusCode: code,
	}
}

/*
ListProjectAWSVPCsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectAWSVPCsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project a w s v p cs default response
func (o *ListProjectAWSVPCsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project a w s v p cs default response has a 2xx status code
func (o *ListProjectAWSVPCsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project a w s v p cs default response has a 3xx status code
func (o *ListProjectAWSVPCsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project a w s v p cs default response has a 4xx status code
func (o *ListProjectAWSVPCsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project a w s v p cs default response has a 5xx status code
func (o *ListProjectAWSVPCsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project a w s v p cs default response a status code equal to that given
func (o *ListProjectAWSVPCsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectAWSVPCsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/vpcs][%d] listProjectAWSVPCs default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectAWSVPCsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/aws/{dc}/vpcs][%d] listProjectAWSVPCs default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectAWSVPCsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectAWSVPCsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}