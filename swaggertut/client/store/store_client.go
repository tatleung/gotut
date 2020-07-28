// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new store API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for store API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteOrder(params *DeleteOrderParams) error

	GetInventory(params *GetInventoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetInventoryOK, error)

	GetOrderByID(params *GetOrderByIDParams) (*GetOrderByIDOK, error)

	PlaceOrder(params *PlaceOrderParams) (*PlaceOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteOrder deletes purchase order by ID

  For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors
*/
func (a *Client) DeleteOrder(params *DeleteOrderParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrderParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrder",
		Method:             "DELETE",
		PathPattern:        "/store/order/{orderId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  GetInventory returns pet inventories by status

  Returns a map of status codes to quantities
*/
func (a *Client) GetInventory(params *GetInventoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetInventoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInventory",
		Method:             "GET",
		PathPattern:        "/store/inventory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInventoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInventoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInventory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrderByID finds purchase order by ID

  For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
*/
func (a *Client) GetOrderByID(params *GetOrderByIDParams) (*GetOrderByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrderById",
		Method:             "GET",
		PathPattern:        "/store/order/{orderId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrderByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrderByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrderById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PlaceOrder places an order for a pet
*/
func (a *Client) PlaceOrder(params *PlaceOrderParams) (*PlaceOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlaceOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "placeOrder",
		Method:             "POST",
		PathPattern:        "/store/order",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PlaceOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PlaceOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for placeOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}