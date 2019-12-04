// Code generated by github.com/swaggest/swac v0.1.8, DO NOT EDIT.

// Package petstore contains autogenerated REST client.
package petstore

import (
	"context"
	"net/http"
	"time"
)

// DefaultBaseURL is the default base URL for this service.
const DefaultBaseURL = "http://petstore.swagger.io/api"

// Client is a REST service HTTP client.
// Swagger Petstore.
// A sample API that uses a petstore as an example to demonstrate features in the swagger-2.0 specification
//
// Version: 1.0.0.
//
// License: Apache 2.0 https://www.apache.org/licenses/LICENSE-2.0.html.
//
// Terms of service: http://swagger.io/terms/.
//
// Contact: Swagger API Team apiteam@swagger.io http://swagger.io.
type Client struct {
	BaseURL           string
	Timeout           time.Duration
	// InstrumentCtxFunc allows adding operation info to context.
	// A pointer to request structure passed into the function.
	// Nil value is ignored.
	InstrumentCtxFunc func(ctx context.Context, method, pattern string, reqStruct interface{}) context.Context
	transport         http.RoundTripper
}

// NewClient creates client instance with default transport.
func NewClient() *Client {
	return &Client{
		transport: http.DefaultTransport,
		Timeout:   30 * time.Second,
		BaseURL:   DefaultBaseURL,
	}
}

// SetTransport sets client transport.
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.transport = transport
}
