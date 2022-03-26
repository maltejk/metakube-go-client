// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// MigrateClusterToExternalCCMReader is a Reader for the MigrateClusterToExternalCCM structure.
type MigrateClusterToExternalCCMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateClusterToExternalCCMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateClusterToExternalCCMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewMigrateClusterToExternalCCMUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMigrateClusterToExternalCCMForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewMigrateClusterToExternalCCMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMigrateClusterToExternalCCMOK creates a MigrateClusterToExternalCCMOK with default headers values
func NewMigrateClusterToExternalCCMOK() *MigrateClusterToExternalCCMOK {
	return &MigrateClusterToExternalCCMOK{}
}

/* MigrateClusterToExternalCCMOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type MigrateClusterToExternalCCMOK struct {
}

func (o *MigrateClusterToExternalCCMOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMOK ", 200)
}

func (o *MigrateClusterToExternalCCMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateClusterToExternalCCMUnauthorized creates a MigrateClusterToExternalCCMUnauthorized with default headers values
func NewMigrateClusterToExternalCCMUnauthorized() *MigrateClusterToExternalCCMUnauthorized {
	return &MigrateClusterToExternalCCMUnauthorized{}
}

/* MigrateClusterToExternalCCMUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type MigrateClusterToExternalCCMUnauthorized struct {
}

func (o *MigrateClusterToExternalCCMUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMUnauthorized ", 401)
}

func (o *MigrateClusterToExternalCCMUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateClusterToExternalCCMForbidden creates a MigrateClusterToExternalCCMForbidden with default headers values
func NewMigrateClusterToExternalCCMForbidden() *MigrateClusterToExternalCCMForbidden {
	return &MigrateClusterToExternalCCMForbidden{}
}

/* MigrateClusterToExternalCCMForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type MigrateClusterToExternalCCMForbidden struct {
}

func (o *MigrateClusterToExternalCCMForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCMForbidden ", 403)
}

func (o *MigrateClusterToExternalCCMForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateClusterToExternalCCMDefault creates a MigrateClusterToExternalCCMDefault with default headers values
func NewMigrateClusterToExternalCCMDefault(code int) *MigrateClusterToExternalCCMDefault {
	return &MigrateClusterToExternalCCMDefault{
		_statusCode: code,
	}
}

/* MigrateClusterToExternalCCMDefault describes a response with status code -1, with default header values.

errorResponse
*/
type MigrateClusterToExternalCCMDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the migrate cluster to external c c m default response
func (o *MigrateClusterToExternalCCMDefault) Code() int {
	return o._statusCode
}

func (o *MigrateClusterToExternalCCMDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration][%d] migrateClusterToExternalCCM default  %+v", o._statusCode, o.Payload)
}
func (o *MigrateClusterToExternalCCMDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MigrateClusterToExternalCCMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
