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

// GetNodeDeploymentRequestReader is a Reader for the GetNodeDeploymentRequest structure.
type GetNodeDeploymentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeDeploymentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeDeploymentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNodeDeploymentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeDeploymentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNodeDeploymentRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeDeploymentRequestOK creates a GetNodeDeploymentRequestOK with default headers values
func NewGetNodeDeploymentRequestOK() *GetNodeDeploymentRequestOK {
	return &GetNodeDeploymentRequestOK{}
}

/* GetNodeDeploymentRequestOK describes a response with status code 200, with default header values.

NodeDeploymentRequest
*/
type GetNodeDeploymentRequestOK struct {
	Payload *models.NodeDeploymentRequest
}

func (o *GetNodeDeploymentRequestOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}][%d] getNodeDeploymentRequestOK  %+v", 200, o.Payload)
}
func (o *GetNodeDeploymentRequestOK) GetPayload() *models.NodeDeploymentRequest {
	return o.Payload
}

func (o *GetNodeDeploymentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeploymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDeploymentRequestUnauthorized creates a GetNodeDeploymentRequestUnauthorized with default headers values
func NewGetNodeDeploymentRequestUnauthorized() *GetNodeDeploymentRequestUnauthorized {
	return &GetNodeDeploymentRequestUnauthorized{}
}

/* GetNodeDeploymentRequestUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetNodeDeploymentRequestUnauthorized struct {
}

func (o *GetNodeDeploymentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}][%d] getNodeDeploymentRequestUnauthorized ", 401)
}

func (o *GetNodeDeploymentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeDeploymentRequestForbidden creates a GetNodeDeploymentRequestForbidden with default headers values
func NewGetNodeDeploymentRequestForbidden() *GetNodeDeploymentRequestForbidden {
	return &GetNodeDeploymentRequestForbidden{}
}

/* GetNodeDeploymentRequestForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetNodeDeploymentRequestForbidden struct {
}

func (o *GetNodeDeploymentRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}][%d] getNodeDeploymentRequestForbidden ", 403)
}

func (o *GetNodeDeploymentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeDeploymentRequestDefault creates a GetNodeDeploymentRequestDefault with default headers values
func NewGetNodeDeploymentRequestDefault(code int) *GetNodeDeploymentRequestDefault {
	return &GetNodeDeploymentRequestDefault{
		_statusCode: code,
	}
}

/* GetNodeDeploymentRequestDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetNodeDeploymentRequestDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get node deployment request default response
func (o *GetNodeDeploymentRequestDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeDeploymentRequestDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}][%d] getNodeDeploymentRequest default  %+v", o._statusCode, o.Payload)
}
func (o *GetNodeDeploymentRequestDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNodeDeploymentRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
