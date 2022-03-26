// Code generated by go-swagger; DO NOT EDIT.

package alibaba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/maltejk/metakube-go-client/pkg/models"
)

// ListAlibabaVSwitchesNoCredentialsV2Reader is a Reader for the ListAlibabaVSwitchesNoCredentialsV2 structure.
type ListAlibabaVSwitchesNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAlibabaVSwitchesNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAlibabaVSwitchesNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAlibabaVSwitchesNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAlibabaVSwitchesNoCredentialsV2OK creates a ListAlibabaVSwitchesNoCredentialsV2OK with default headers values
func NewListAlibabaVSwitchesNoCredentialsV2OK() *ListAlibabaVSwitchesNoCredentialsV2OK {
	return &ListAlibabaVSwitchesNoCredentialsV2OK{}
}

/* ListAlibabaVSwitchesNoCredentialsV2OK describes a response with status code 200, with default header values.

AlibabaVSwitchList
*/
type ListAlibabaVSwitchesNoCredentialsV2OK struct {
	Payload models.AlibabaVSwitchList
}

func (o *ListAlibabaVSwitchesNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/vswitches][%d] listAlibabaVSwitchesNoCredentialsV2OK  %+v", 200, o.Payload)
}
func (o *ListAlibabaVSwitchesNoCredentialsV2OK) GetPayload() models.AlibabaVSwitchList {
	return o.Payload
}

func (o *ListAlibabaVSwitchesNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlibabaVSwitchesNoCredentialsV2Default creates a ListAlibabaVSwitchesNoCredentialsV2Default with default headers values
func NewListAlibabaVSwitchesNoCredentialsV2Default(code int) *ListAlibabaVSwitchesNoCredentialsV2Default {
	return &ListAlibabaVSwitchesNoCredentialsV2Default{
		_statusCode: code,
	}
}

/* ListAlibabaVSwitchesNoCredentialsV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type ListAlibabaVSwitchesNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list alibaba v switches no credentials v2 default response
func (o *ListAlibabaVSwitchesNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListAlibabaVSwitchesNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/vswitches][%d] listAlibabaVSwitchesNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}
func (o *ListAlibabaVSwitchesNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAlibabaVSwitchesNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}