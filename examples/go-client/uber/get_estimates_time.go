// Code generated by github.com/swaggest/swac v0.1.8, DO NOT EDIT.

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

// GetEstimatesTimeRequest is operation request value.
type GetEstimatesTimeRequest struct {
	// StartLatitude is a required `start_latitude` parameter in query.
	// Latitude component of start location.
	StartLatitude  float64
	// StartLongitude is a required `start_longitude` parameter in query.
	// Longitude component of start location.
	StartLongitude float64
	// CustomerUuid is an optional `customer_uuid` parameter in query.
	// Unique customer identifier to be used for experience customization.
	CustomerUuid   *string
	// ProductID is an optional `product_id` parameter in query.
	// Unique identifier representing a specific product for a given latitude & longitude.
	ProductID      *string
}

// encode creates *http.Request for request data.
func (request *GetEstimatesTimeRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/estimates/time"

	query := make(url.Values, 4)
	query.Set("start_latitude", strconv.FormatFloat(request.StartLatitude, 'g', -1, 64))

	query.Set("start_longitude", strconv.FormatFloat(request.StartLongitude, 'g', -1, 64))

	if request.CustomerUuid != nil {
		query.Set("customer_uuid", *request.CustomerUuid)
	}

	if request.ProductID != nil {
		query.Set("product_id", *request.ProductID)
	}

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

// GetEstimatesTimeResponse is operation response value.
type GetEstimatesTimeResponse struct {
	StatusCode int
	ValueOK    []Product  // ValueOK is a value of 200 OK response.
	Default    *Error     // Default is a default value of response.
}

// decode loads data from *http.Response.
func (result *GetEstimatesTimeResponse) decode(resp *http.Response) error {
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

// GetEstimatesTime performs REST operation.
func (c *Client) GetEstimatesTime(ctx context.Context, request GetEstimatesTimeRequest) (result GetEstimatesTimeResponse, err error) {
	if c.InstrumentCtxFunc != nil {
		ctx = c.InstrumentCtxFunc(ctx, http.MethodGet, "/estimates/time", &request)
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
