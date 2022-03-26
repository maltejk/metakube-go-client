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

// ListExternalClusterNodesReader is a Reader for the ListExternalClusterNodes structure.
type ListExternalClusterNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExternalClusterNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExternalClusterNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListExternalClusterNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListExternalClusterNodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListExternalClusterNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListExternalClusterNodesOK creates a ListExternalClusterNodesOK with default headers values
func NewListExternalClusterNodesOK() *ListExternalClusterNodesOK {
	return &ListExternalClusterNodesOK{}
}

/* ListExternalClusterNodesOK describes a response with status code 200, with default header values.

Node
*/
type ListExternalClusterNodesOK struct {
	Payload []*models.Node
}

func (o *ListExternalClusterNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesOK  %+v", 200, o.Payload)
}
func (o *ListExternalClusterNodesOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *ListExternalClusterNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListExternalClusterNodesUnauthorized creates a ListExternalClusterNodesUnauthorized with default headers values
func NewListExternalClusterNodesUnauthorized() *ListExternalClusterNodesUnauthorized {
	return &ListExternalClusterNodesUnauthorized{}
}

/* ListExternalClusterNodesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterNodesUnauthorized struct {
}

func (o *ListExternalClusterNodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesUnauthorized ", 401)
}

func (o *ListExternalClusterNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterNodesForbidden creates a ListExternalClusterNodesForbidden with default headers values
func NewListExternalClusterNodesForbidden() *ListExternalClusterNodesForbidden {
	return &ListExternalClusterNodesForbidden{}
}

/* ListExternalClusterNodesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListExternalClusterNodesForbidden struct {
}

func (o *ListExternalClusterNodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodesForbidden ", 403)
}

func (o *ListExternalClusterNodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListExternalClusterNodesDefault creates a ListExternalClusterNodesDefault with default headers values
func NewListExternalClusterNodesDefault(code int) *ListExternalClusterNodesDefault {
	return &ListExternalClusterNodesDefault{
		_statusCode: code,
	}
}

/* ListExternalClusterNodesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListExternalClusterNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list external cluster nodes default response
func (o *ListExternalClusterNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListExternalClusterNodesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes][%d] listExternalClusterNodes default  %+v", o._statusCode, o.Payload)
}
func (o *ListExternalClusterNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListExternalClusterNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
