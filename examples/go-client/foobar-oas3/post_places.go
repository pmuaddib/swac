// Code generated by github.com/swaggest/swac <version>, DO NOT EDIT.

package foobar

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// PostPlacesRequest is operation request value.
type PostPlacesRequest struct {
	Body *UsecaseCreatePlaceInput  // Body is a JSON request body.
}

// encode creates *http.Request for request data.
func (request *PostPlacesRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/places"

	body, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, requestURI, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	req = req.WithContext(ctx)

	return req, err
}

// PostPlacesResponse is operation response value.
type PostPlacesResponse struct {
	StatusCode               int
	RawBody                  []byte            // RawBody contains read bytes of response body.
	ValueOK                  *PlaceEntity      // ValueOK is a value of 200 OK response.
	ValueBadRequest          *RestErrResponse  // ValueBadRequest is a value of 400 Bad Request response.
	ValueConflict            *RestErrResponse  // ValueConflict is a value of 409 Conflict response.
	ValueInternalServerError *RestErrResponse  // ValueInternalServerError is a value of 500 Internal Server Error response.
}

// decode loads data from *http.Response.
func (result *PostPlacesResponse) decode(resp *http.Response) error {
	var err error

	dump := bytes.NewBuffer(nil)
	body := io.TeeReader(resp.Body, dump)

	result.StatusCode = resp.StatusCode

	switch resp.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(body).Decode(&result.ValueOK)
	case http.StatusBadRequest:
		err = json.NewDecoder(body).Decode(&result.ValueBadRequest)
	case http.StatusConflict:
		err = json.NewDecoder(body).Decode(&result.ValueConflict)
	case http.StatusInternalServerError:
		err = json.NewDecoder(body).Decode(&result.ValueInternalServerError)
	default:
		_, readErr := ioutil.ReadAll(body)
		if readErr != nil {
			err = errors.New("unexpected response status: " + resp.Status +
				", could not read response body: " + readErr.Error())
		} else {
			err = errors.New("unexpected response status: " + resp.Status)
		}
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

// PostPlaces performs REST operation.
func (c *Client) PostPlaces(ctx context.Context, request PostPlacesRequest) (result PostPlacesResponse, err error) {
	if c.InstrumentCtxFunc != nil {
		ctx = c.InstrumentCtxFunc(ctx, http.MethodPost, "/places", &request)
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
