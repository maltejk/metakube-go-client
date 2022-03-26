// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListCredentialsReader is a Reader for the ListCredentials structure.
type ListCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCredentialsOK creates a ListCredentialsOK with default headers values
func NewListCredentialsOK() *ListCredentialsOK {
	return &ListCredentialsOK{}
}

/* ListCredentialsOK describes a response with status code 200, with default header values.

CredentialList
*/
type ListCredentialsOK struct {
	Payload *models.CredentialList
}

func (o *ListCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/presets/credentials][%d] listCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListCredentialsOK) GetPayload() *models.CredentialList {
	return o.Payload
}

func (o *ListCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCredentialsDefault creates a ListCredentialsDefault with default headers values
func NewListCredentialsDefault(code int) *ListCredentialsDefault {
	return &ListCredentialsDefault{
		_statusCode: code,
	}
}

/* ListCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list credentials default response
func (o *ListCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/presets/credentials][%d] listCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
