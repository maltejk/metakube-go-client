// Code generated by go-swagger; DO NOT EDIT.

package allowedregistry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListAllowedRegistriesReader is a Reader for the ListAllowedRegistries structure.
type ListAllowedRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllowedRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllowedRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAllowedRegistriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAllowedRegistriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAllowedRegistriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAllowedRegistriesOK creates a ListAllowedRegistriesOK with default headers values
func NewListAllowedRegistriesOK() *ListAllowedRegistriesOK {
	return &ListAllowedRegistriesOK{}
}

/* ListAllowedRegistriesOK describes a response with status code 200, with default header values.

AllowedRegistry
*/
type ListAllowedRegistriesOK struct {
	Payload []*models.AllowedRegistry
}

func (o *ListAllowedRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesOK  %+v", 200, o.Payload)
}
func (o *ListAllowedRegistriesOK) GetPayload() []*models.AllowedRegistry {
	return o.Payload
}

func (o *ListAllowedRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllowedRegistriesUnauthorized creates a ListAllowedRegistriesUnauthorized with default headers values
func NewListAllowedRegistriesUnauthorized() *ListAllowedRegistriesUnauthorized {
	return &ListAllowedRegistriesUnauthorized{}
}

/* ListAllowedRegistriesUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListAllowedRegistriesUnauthorized struct {
}

func (o *ListAllowedRegistriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesUnauthorized ", 401)
}

func (o *ListAllowedRegistriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllowedRegistriesForbidden creates a ListAllowedRegistriesForbidden with default headers values
func NewListAllowedRegistriesForbidden() *ListAllowedRegistriesForbidden {
	return &ListAllowedRegistriesForbidden{}
}

/* ListAllowedRegistriesForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListAllowedRegistriesForbidden struct {
}

func (o *ListAllowedRegistriesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistriesForbidden ", 403)
}

func (o *ListAllowedRegistriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllowedRegistriesDefault creates a ListAllowedRegistriesDefault with default headers values
func NewListAllowedRegistriesDefault(code int) *ListAllowedRegistriesDefault {
	return &ListAllowedRegistriesDefault{
		_statusCode: code,
	}
}

/* ListAllowedRegistriesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAllowedRegistriesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list allowed registries default response
func (o *ListAllowedRegistriesDefault) Code() int {
	return o._statusCode
}

func (o *ListAllowedRegistriesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/allowedregistries][%d] listAllowedRegistries default  %+v", o._statusCode, o.Payload)
}
func (o *ListAllowedRegistriesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAllowedRegistriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}