// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListRoleBindingReader is a Reader for the ListRoleBinding structure.
type ListRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRoleBindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRoleBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRoleBindingOK creates a ListRoleBindingOK with default headers values
func NewListRoleBindingOK() *ListRoleBindingOK {
	return &ListRoleBindingOK{}
}

/* ListRoleBindingOK describes a response with status code 200, with default header values.

RoleBinding
*/
type ListRoleBindingOK struct {
	Payload []*models.RoleBinding
}

func (o *ListRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/bindings][%d] listRoleBindingOK  %+v", 200, o.Payload)
}
func (o *ListRoleBindingOK) GetPayload() []*models.RoleBinding {
	return o.Payload
}

func (o *ListRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRoleBindingUnauthorized creates a ListRoleBindingUnauthorized with default headers values
func NewListRoleBindingUnauthorized() *ListRoleBindingUnauthorized {
	return &ListRoleBindingUnauthorized{}
}

/* ListRoleBindingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListRoleBindingUnauthorized struct {
}

func (o *ListRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/bindings][%d] listRoleBindingUnauthorized ", 401)
}

func (o *ListRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleBindingForbidden creates a ListRoleBindingForbidden with default headers values
func NewListRoleBindingForbidden() *ListRoleBindingForbidden {
	return &ListRoleBindingForbidden{}
}

/* ListRoleBindingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListRoleBindingForbidden struct {
}

func (o *ListRoleBindingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/bindings][%d] listRoleBindingForbidden ", 403)
}

func (o *ListRoleBindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleBindingDefault creates a ListRoleBindingDefault with default headers values
func NewListRoleBindingDefault(code int) *ListRoleBindingDefault {
	return &ListRoleBindingDefault{
		_statusCode: code,
	}
}

/* ListRoleBindingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListRoleBindingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list role binding default response
func (o *ListRoleBindingDefault) Code() int {
	return o._statusCode
}

func (o *ListRoleBindingDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/bindings][%d] listRoleBinding default  %+v", o._statusCode, o.Payload)
}
func (o *ListRoleBindingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListRoleBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}