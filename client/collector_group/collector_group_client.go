// Code generated by go-swagger; DO NOT EDIT.

package collector_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new collector group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for collector group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
AddCollectorGroup adds collector group
*/
func (a *Client) AddCollectorGroup(params *AddCollectorGroupParams) (*AddCollectorGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCollectorGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCollectorGroup",
		Method:             "POST",
		PathPattern:        "/setting/collector/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddCollectorGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddCollectorGroupOK), nil

}

/*
DeleteCollectorGroupByID deletes collector group
*/
func (a *Client) DeleteCollectorGroupByID(params *DeleteCollectorGroupByIDParams) (*DeleteCollectorGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCollectorGroupByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCollectorGroupById",
		Method:             "DELETE",
		PathPattern:        "/setting/collector/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCollectorGroupByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCollectorGroupByIDOK), nil

}

/*
GetCollectorGroupByID gets collector group
*/
func (a *Client) GetCollectorGroupByID(params *GetCollectorGroupByIDParams) (*GetCollectorGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectorGroupByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCollectorGroupById",
		Method:             "GET",
		PathPattern:        "/setting/collector/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectorGroupByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCollectorGroupByIDOK), nil

}

/*
GetCollectorGroupList gets collector group list
*/
func (a *Client) GetCollectorGroupList(params *GetCollectorGroupListParams) (*GetCollectorGroupListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectorGroupListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCollectorGroupList",
		Method:             "GET",
		PathPattern:        "/setting/collector/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectorGroupListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCollectorGroupListOK), nil

}

/*
UpdateCollectorGroupByID updates collector group
*/
func (a *Client) UpdateCollectorGroupByID(params *UpdateCollectorGroupByIDParams) (*UpdateCollectorGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectorGroupByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCollectorGroupById",
		Method:             "PUT",
		PathPattern:        "/setting/collector/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCollectorGroupByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateCollectorGroupByIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
