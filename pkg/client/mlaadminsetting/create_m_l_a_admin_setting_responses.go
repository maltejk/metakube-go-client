// Code generated by go-swagger; DO NOT EDIT.

package mlaadminsetting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// CreateMLAAdminSettingReader is a Reader for the CreateMLAAdminSetting structure.
type CreateMLAAdminSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMLAAdminSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMLAAdminSettingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateMLAAdminSettingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMLAAdminSettingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateMLAAdminSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMLAAdminSettingCreated creates a CreateMLAAdminSettingCreated with default headers values
func NewCreateMLAAdminSettingCreated() *CreateMLAAdminSettingCreated {
	return &CreateMLAAdminSettingCreated{}
}

/* CreateMLAAdminSettingCreated describes a response with status code 201, with default header values.

MLAAdminSetting
*/
type CreateMLAAdminSettingCreated struct {
	Payload *models.MLAAdminSetting
}

func (o *CreateMLAAdminSettingCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] createMLAAdminSettingCreated  %+v", 201, o.Payload)
}
func (o *CreateMLAAdminSettingCreated) GetPayload() *models.MLAAdminSetting {
	return o.Payload
}

func (o *CreateMLAAdminSettingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MLAAdminSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMLAAdminSettingUnauthorized creates a CreateMLAAdminSettingUnauthorized with default headers values
func NewCreateMLAAdminSettingUnauthorized() *CreateMLAAdminSettingUnauthorized {
	return &CreateMLAAdminSettingUnauthorized{}
}

/* CreateMLAAdminSettingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateMLAAdminSettingUnauthorized struct {
}

func (o *CreateMLAAdminSettingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] createMLAAdminSettingUnauthorized ", 401)
}

func (o *CreateMLAAdminSettingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMLAAdminSettingForbidden creates a CreateMLAAdminSettingForbidden with default headers values
func NewCreateMLAAdminSettingForbidden() *CreateMLAAdminSettingForbidden {
	return &CreateMLAAdminSettingForbidden{}
}

/* CreateMLAAdminSettingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateMLAAdminSettingForbidden struct {
}

func (o *CreateMLAAdminSettingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] createMLAAdminSettingForbidden ", 403)
}

func (o *CreateMLAAdminSettingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMLAAdminSettingDefault creates a CreateMLAAdminSettingDefault with default headers values
func NewCreateMLAAdminSettingDefault(code int) *CreateMLAAdminSettingDefault {
	return &CreateMLAAdminSettingDefault{
		_statusCode: code,
	}
}

/* CreateMLAAdminSettingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateMLAAdminSettingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create m l a admin setting default response
func (o *CreateMLAAdminSettingDefault) Code() int {
	return o._statusCode
}

func (o *CreateMLAAdminSettingDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] createMLAAdminSetting default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMLAAdminSettingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateMLAAdminSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}