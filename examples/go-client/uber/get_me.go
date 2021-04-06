// Code generated by github.com/swaggest/swac <version>, DO NOT EDIT.

package uber

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// GetMeRequest is operation request value.
type GetMeRequest struct {

}

// encode creates *http.Request for request data.
func (request *GetMeRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/me"

	req, err := http.NewRequest(http.MethodGet, requestURI, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	req = req.WithContext(ctx)

	return req, err
}

// GetMeResponse is operation response value.
type GetMeResponse struct {
	StatusCode int
	RawBody    []byte    // RawBody contains read bytes of response body.
	ValueOK    *Profile  // ValueOK is a value of 200 OK response.
	Default    *Error    // Default is a default value of response.
}

// decode loads data from *http.Response.
func (result *GetMeResponse) decode(resp *http.Response) error {
	var err error

	dump := bytes.NewBuffer(nil)
	body := io.TeeReader(resp.Body, dump)

	result.StatusCode = resp.StatusCode

	switch resp.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(body).Decode(&result.ValueOK)
	default:
		err = json.NewDecoder(body).Decode(&result.Default)
	}

	result.RawBody = dump.Bytes()

	if err != nil {
		return responseError{
			resp: resp,
			body: dump.Bytes(),
			err:  err,
		}
	}

	return nil
}

// GetMe performs REST operation.
func (c *Client) GetMe(ctx context.Context, request GetMeRequest) (result GetMeResponse, err error) {
	if c.InstrumentCtxFunc != nil {
		ctx = c.InstrumentCtxFunc(ctx, http.MethodGet, "/me", &request)
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
