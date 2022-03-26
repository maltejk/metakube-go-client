// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListAzureAvailabilityZonesNoCredentialsV2Reader is a Reader for the ListAzureAvailabilityZonesNoCredentialsV2 structure.
type ListAzureAvailabilityZonesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureAvailabilityZonesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureAvailabilityZonesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureAvailabilityZonesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureAvailabilityZonesNoCredentialsV2OK creates a ListAzureAvailabilityZonesNoCredentialsV2OK with default headers values
func NewListAzureAvailabilityZonesNoCredentialsV2OK() *ListAzureAvailabilityZonesNoCredentialsV2OK {
	return &ListAzureAvailabilityZonesNoCredentialsV2OK{}
}

/* ListAzureAvailabilityZonesNoCredentialsV2OK describes a response with status code 200, with default header values.

AzureAvailabilityZonesList
*/
type ListAzureAvailabilityZonesNoCredentialsV2OK struct {
	Payload *models.AzureAvailabilityZonesList
}

func (o *ListAzureAvailabilityZonesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/availabilityzones][%d] listAzureAvailabilityZonesNoCredentialsV2OK  %+v", 200, o.Payload)
}
func (o *ListAzureAvailabilityZonesNoCredentialsV2OK) GetPayload() *models.AzureAvailabilityZonesList {
	return o.Payload
}

func (o *ListAzureAvailabilityZonesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureAvailabilityZonesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureAvailabilityZonesNoCredentialsV2Default creates a ListAzureAvailabilityZonesNoCredentialsV2Default with default headers values
func NewListAzureAvailabilityZonesNoCredentialsV2Default(code int) *ListAzureAvailabilityZonesNoCredentialsV2Default {
	return &ListAzureAvailabilityZonesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListAzureAvailabilityZonesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureAvailabilityZonesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure availability zones no credentials v2 default response
func (o *ListAzureAvailabilityZonesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAzureAvailabilityZonesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/availabilityzones][%d] listAzureAvailabilityZonesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}
func (o *ListAzureAvailabilityZonesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureAvailabilityZonesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
