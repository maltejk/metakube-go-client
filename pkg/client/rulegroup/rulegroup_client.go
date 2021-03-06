// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rulegroup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rulegroup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRuleGroup(params *CreateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRuleGroupCreated, error)

	DeleteRuleGroup(params *DeleteRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRuleGroupOK, error)

	GetRuleGroup(params *GetRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuleGroupOK, error)

	ListRuleGroups(params *ListRuleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRuleGroupsOK, error)

	UpdateRuleGroup(params *UpdateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRuleGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRuleGroup Creates a rule group that will belong to the given cluster
*/
func (a *Client) CreateRuleGroup(params *CreateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRuleGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRuleGroup",
		Method:             "POST",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRuleGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRuleGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRuleGroup deletes the given rule group that belongs to the cluster
*/
func (a *Client) DeleteRuleGroup(params *DeleteRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRuleGroup",
		Method:             "DELETE",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRuleGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRuleGroup gets a specified rule group for the given cluster
*/
func (a *Client) GetRuleGroup(params *GetRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRuleGroup",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRuleGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListRuleGroups Lists rule groups that belong to the given cluster
*/
func (a *Client) ListRuleGroups(params *ListRuleGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRuleGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRuleGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRuleGroups",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRuleGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRuleGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRuleGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateRuleGroup updates the specified rule group for the given cluster
*/
func (a *Client) UpdateRuleGroup(params *UpdateRuleGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRuleGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRuleGroup",
		Method:             "PUT",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRuleGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRuleGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRuleGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
