// Code generated by github.com/swaggest/swac v0.1.5, DO NOT EDIT.

package uber

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// GetProductsRequest is operation request value.
type GetProductsRequest struct {
	// Latitude is a required `latitude` parameter in query.
	// Latitude component of location.
	Latitude  float64
	// Longitude is a required `longitude` parameter in query.
	// Longitude component of location.
	Longitude float64
}

// encode creates *http.Request for request data.
func (request *GetProductsRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/products"

	query := make(url.Values, 2)
	query.Set("latitude", strconv.FormatFloat(request.Latitude, 'g', -1, 64))

	query.Set("longitude", strconv.FormatFloat(request.Longitude, 'g', -1, 64))

	if len(query) > 0 {
		requestURI += "?" + query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestURI, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	req = req.WithContext(ctx)

	return req, err
}

// GetProductsResponse is operation response value.
type GetProductsResponse struct {
	StatusCode int
	ValueOK    []Product  // ValueOK is a value of 200 OK response.
	Default    *Error     // Default is a default value of response.
}

// decode loads data from *http.Response.
func (result *GetProductsResponse) decode(resp *http.Response) error {
	var err error

	dump := bytes.NewBuffer(nil)
	body := io.TeeReader(resp.Body, dump)

	result.StatusCode = resp.StatusCode

	switch resp.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(body).Decode(&result.ValueOK)
	default:
		err = json.NewDecoder(resp.Body).Decode(&result.Default)
	}

	if err != nil {
		return responseError{
			resp: resp,
			body: dump.Bytes(),
			err:  err,
		}
	}

	return nil
}

// GetProducts performs REST operation.
func (c *Client) GetProducts(ctx context.Context, request GetProductsRequest) (result GetProductsResponse, err error) {
	if c.InstrumentCtxFunc != nil {
		ctx = c.InstrumentCtxFunc(ctx, http.MethodGet, "/products", &request)
	}

	if c.Timeout != 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, c.Timeout)

		defer cancel()
	}

	req, err := request.encode(ctx, c.BaseURL)
	if err != nil {
		return result, err
	}

	resp, err := c.transport.RoundTrip(req)

	if err != nil {
		return result, err
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	err = result.decode(resp)

	return result, err
}
