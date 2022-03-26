// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListOpenstackSecurityGroupsNoCredentialsReader is a Reader for the ListOpenstackSecurityGroupsNoCredentials structure.
type ListOpenstackSecurityGroupsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackSecurityGroupsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackSecurityGroupsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackSecurityGroupsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackSecurityGroupsNoCredentialsOK creates a ListOpenstackSecurityGroupsNoCredentialsOK with default headers values
func NewListOpenstackSecurityGroupsNoCredentialsOK() *ListOpenstackSecurityGroupsNoCredentialsOK {
	return &ListOpenstackSecurityGroupsNoCredentialsOK{}
}

/* ListOpenstackSecurityGroupsNoCredentialsOK describes a response with status code 200, with default header values.

OpenstackSecurityGroup
*/
type ListOpenstackSecurityGroupsNoCredentialsOK struct {
	Payload []*models.OpenstackSecurityGroup
}

func (o *ListOpenstackSecurityGroupsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/securitygroups][%d] listOpenstackSecurityGroupsNoCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListOpenstackSecurityGroupsNoCredentialsOK) GetPayload() []*models.OpenstackSecurityGroup {
	return o.Payload
}

func (o *ListOpenstackSecurityGroupsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackSecurityGroupsNoCredentialsDefault creates a ListOpenstackSecurityGroupsNoCredentialsDefault with default headers values
func NewListOpenstackSecurityGroupsNoCredentialsDefault(code int) *ListOpenstackSecurityGroupsNoCredentialsDefault {
	return &ListOpenstackSecurityGroupsNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListOpenstackSecurityGroupsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackSecurityGroupsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack security groups no credentials default response
func (o *ListOpenstackSecurityGroupsNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackSecurityGroupsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/securitygroups][%d] listOpenstackSecurityGroupsNoCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListOpenstackSecurityGroupsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackSecurityGroupsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
