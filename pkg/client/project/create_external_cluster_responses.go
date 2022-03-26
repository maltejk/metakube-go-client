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

// CreateExternalClusterReader is a Reader for the CreateExternalCluster structure.
type CreateExternalClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExternalClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateExternalClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateExternalClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateExternalClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateExternalClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateExternalClusterCreated creates a CreateExternalClusterCreated with default headers values
func NewCreateExternalClusterCreated() *CreateExternalClusterCreated {
	return &CreateExternalClusterCreated{}
}

/* CreateExternalClusterCreated describes a response with status code 201, with default header values.

Cluster
*/
type CreateExternalClusterCreated struct {
	Payload *models.Cluster
}

func (o *CreateExternalClusterCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters][%d] createExternalClusterCreated  %+v", 201, o.Payload)
}
func (o *CreateExternalClusterCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *CreateExternalClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExternalClusterUnauthorized creates a CreateExternalClusterUnauthorized with default headers values
func NewCreateExternalClusterUnauthorized() *CreateExternalClusterUnauthorized {
	return &CreateExternalClusterUnauthorized{}
}

/* CreateExternalClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateExternalClusterUnauthorized struct {
}

func (o *CreateExternalClusterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters][%d] createExternalClusterUnauthorized ", 401)
}

func (o *CreateExternalClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExternalClusterForbidden creates a CreateExternalClusterForbidden with default headers values
func NewCreateExternalClusterForbidden() *CreateExternalClusterForbidden {
	return &CreateExternalClusterForbidden{}
}

/* CreateExternalClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateExternalClusterForbidden struct {
}

func (o *CreateExternalClusterForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters][%d] createExternalClusterForbidden ", 403)
}

func (o *CreateExternalClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExternalClusterDefault creates a CreateExternalClusterDefault with default headers values
func NewCreateExternalClusterDefault(code int) *CreateExternalClusterDefault {
	return &CreateExternalClusterDefault{
		_statusCode: code,
	}
}

/* CreateExternalClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateExternalClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create external cluster default response
func (o *CreateExternalClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreateExternalClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/kubernetes/clusters][%d] createExternalCluster default  %+v", o._statusCode, o.Payload)
}
func (o *CreateExternalClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateExternalClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
