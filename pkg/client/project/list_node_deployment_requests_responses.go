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

// ListNodeDeploymentRequestsReader is a Reader for the ListNodeDeploymentRequests structure.
type ListNodeDeploymentRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodeDeploymentRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodeDeploymentRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNodeDeploymentRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListNodeDeploymentRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNodeDeploymentRequestsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodeDeploymentRequestsOK creates a ListNodeDeploymentRequestsOK with default headers values
func NewListNodeDeploymentRequestsOK() *ListNodeDeploymentRequestsOK {
	return &ListNodeDeploymentRequestsOK{}
}

/* ListNodeDeploymentRequestsOK describes a response with status code 200, with default header values.

NodeDeploymentRequest
*/
type ListNodeDeploymentRequestsOK struct {
	Payload []*models.NodeDeploymentRequest
}

func (o *ListNodeDeploymentRequestsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] listNodeDeploymentRequestsOK  %+v", 200, o.Payload)
}
func (o *ListNodeDeploymentRequestsOK) GetPayload() []*models.NodeDeploymentRequest {
	return o.Payload
}

func (o *ListNodeDeploymentRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodeDeploymentRequestsUnauthorized creates a ListNodeDeploymentRequestsUnauthorized with default headers values
func NewListNodeDeploymentRequestsUnauthorized() *ListNodeDeploymentRequestsUnauthorized {
	return &ListNodeDeploymentRequestsUnauthorized{}
}

/* ListNodeDeploymentRequestsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListNodeDeploymentRequestsUnauthorized struct {
}

func (o *ListNodeDeploymentRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] listNodeDeploymentRequestsUnauthorized ", 401)
}

func (o *ListNodeDeploymentRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodeDeploymentRequestsForbidden creates a ListNodeDeploymentRequestsForbidden with default headers values
func NewListNodeDeploymentRequestsForbidden() *ListNodeDeploymentRequestsForbidden {
	return &ListNodeDeploymentRequestsForbidden{}
}

/* ListNodeDeploymentRequestsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListNodeDeploymentRequestsForbidden struct {
}

func (o *ListNodeDeploymentRequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] listNodeDeploymentRequestsForbidden ", 403)
}

func (o *ListNodeDeploymentRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodeDeploymentRequestsDefault creates a ListNodeDeploymentRequestsDefault with default headers values
func NewListNodeDeploymentRequestsDefault(code int) *ListNodeDeploymentRequestsDefault {
	return &ListNodeDeploymentRequestsDefault{
		_statusCode: code,
	}
}

/* ListNodeDeploymentRequestsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNodeDeploymentRequestsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list node deployment requests default response
func (o *ListNodeDeploymentRequestsDefault) Code() int {
	return o._statusCode
}

func (o *ListNodeDeploymentRequestsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] listNodeDeploymentRequests default  %+v", o._statusCode, o.Payload)
}
func (o *ListNodeDeploymentRequestsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNodeDeploymentRequestsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
