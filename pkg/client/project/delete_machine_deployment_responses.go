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

// DeleteMachineDeploymentReader is a Reader for the DeleteMachineDeployment structure.
type DeleteMachineDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMachineDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMachineDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMachineDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMachineDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMachineDeploymentOK creates a DeleteMachineDeploymentOK with default headers values
func NewDeleteMachineDeploymentOK() *DeleteMachineDeploymentOK {
	return &DeleteMachineDeploymentOK{}
}

/* DeleteMachineDeploymentOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteMachineDeploymentOK struct {
}

func (o *DeleteMachineDeploymentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentOK ", 200)
}

func (o *DeleteMachineDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDeploymentUnauthorized creates a DeleteMachineDeploymentUnauthorized with default headers values
func NewDeleteMachineDeploymentUnauthorized() *DeleteMachineDeploymentUnauthorized {
	return &DeleteMachineDeploymentUnauthorized{}
}

/* DeleteMachineDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteMachineDeploymentUnauthorized struct {
}

func (o *DeleteMachineDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentUnauthorized ", 401)
}

func (o *DeleteMachineDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDeploymentForbidden creates a DeleteMachineDeploymentForbidden with default headers values
func NewDeleteMachineDeploymentForbidden() *DeleteMachineDeploymentForbidden {
	return &DeleteMachineDeploymentForbidden{}
}

/* DeleteMachineDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteMachineDeploymentForbidden struct {
}

func (o *DeleteMachineDeploymentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeploymentForbidden ", 403)
}

func (o *DeleteMachineDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDeploymentDefault creates a DeleteMachineDeploymentDefault with default headers values
func NewDeleteMachineDeploymentDefault(code int) *DeleteMachineDeploymentDefault {
	return &DeleteMachineDeploymentDefault{
		_statusCode: code,
	}
}

/* DeleteMachineDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteMachineDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete machine deployment default response
func (o *DeleteMachineDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMachineDeploymentDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}][%d] deleteMachineDeployment default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMachineDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteMachineDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
