// Code generated by go-swagger; DO NOT EDIT.

package anexia

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListAnexiaVlansNoCredentialsV2Reader is a Reader for the ListAnexiaVlansNoCredentialsV2 structure.
type ListAnexiaVlansNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAnexiaVlansNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAnexiaVlansNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAnexiaVlansNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAnexiaVlansNoCredentialsV2OK creates a ListAnexiaVlansNoCredentialsV2OK with default headers values
func NewListAnexiaVlansNoCredentialsV2OK() *ListAnexiaVlansNoCredentialsV2OK {
	return &ListAnexiaVlansNoCredentialsV2OK{}
}

/* ListAnexiaVlansNoCredentialsV2OK describes a response with status code 200, with default header values.

AnexiaVlanList
*/
type ListAnexiaVlansNoCredentialsV2OK struct {
	Payload models.AnexiaVlanList
}

func (o *ListAnexiaVlansNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/vlans][%d] listAnexiaVlansNoCredentialsV2OK  %+v", 200, o.Payload)
}
func (o *ListAnexiaVlansNoCredentialsV2OK) GetPayload() models.AnexiaVlanList {
	return o.Payload
}

func (o *ListAnexiaVlansNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAnexiaVlansNoCredentialsV2Default creates a ListAnexiaVlansNoCredentialsV2Default with default headers values
func NewListAnexiaVlansNoCredentialsV2Default(code int) *ListAnexiaVlansNoCredentialsV2Default {
	return &ListAnexiaVlansNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListAnexiaVlansNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAnexiaVlansNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list anexia vlans no credentials v2 default response
func (o *ListAnexiaVlansNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAnexiaVlansNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/vlans][%d] listAnexiaVlansNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}
func (o *ListAnexiaVlansNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAnexiaVlansNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
