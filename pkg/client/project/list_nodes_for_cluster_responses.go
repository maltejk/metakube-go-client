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

// ListNodesForClusterReader is a Reader for the ListNodesForCluster structure.
type ListNodesForClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodesForClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodesForClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNodesForClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListNodesForClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNodesForClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodesForClusterOK creates a ListNodesForClusterOK with default headers values
func NewListNodesForClusterOK() *ListNodesForClusterOK {
	return &ListNodesForClusterOK{}
}

/* ListNodesForClusterOK describes a response with status code 200, with default header values.

Node
*/
type ListNodesForClusterOK struct {
	Payload []*models.Node
}

func (o *ListNodesForClusterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes][%d] listNodesForClusterOK  %+v", 200, o.Payload)
}
func (o *ListNodesForClusterOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *ListNodesForClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodesForClusterUnauthorized creates a ListNodesForClusterUnauthorized with default headers values
func NewListNodesForClusterUnauthorized() *ListNodesForClusterUnauthorized {
	return &ListNodesForClusterUnauthorized{}
}

/* ListNodesForClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListNodesForClusterUnauthorized struct {
}

func (o *ListNodesForClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes][%d] listNodesForClusterUnauthorized ", 401)
}

func (o *ListNodesForClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodesForClusterForbidden creates a ListNodesForClusterForbidden with default headers values
func NewListNodesForClusterForbidden() *ListNodesForClusterForbidden {
	return &ListNodesForClusterForbidden{}
}

/* ListNodesForClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListNodesForClusterForbidden struct {
}

func (o *ListNodesForClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes][%d] listNodesForClusterForbidden ", 403)
}

func (o *ListNodesForClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNodesForClusterDefault creates a ListNodesForClusterDefault with default headers values
func NewListNodesForClusterDefault(code int) *ListNodesForClusterDefault {
	return &ListNodesForClusterDefault{
		_statusCode: code,
	}
}

/* ListNodesForClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNodesForClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list nodes for cluster default response
func (o *ListNodesForClusterDefault) Code() int {
	return o._statusCode
}

func (o *ListNodesForClusterDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes][%d] listNodesForCluster default  %+v", o._statusCode, o.Payload)
}
func (o *ListNodesForClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNodesForClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
