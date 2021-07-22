// Code generated by github.com/swaggest/swac <version>, DO NOT EDIT.

package acme

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetChallengeResponseStateKeyPersonUIDRequest is operation request value.
type GetChallengeResponseStateKeyPersonUIDRequest struct {
	StateKey  string  // StateKey is a required `state_key` parameter in path.
	PersonUID string  // PersonUID is a required `person_uid` parameter in path.
}

// encode creates *http.Request for request data.
func (request *GetChallengeResponseStateKeyPersonUIDRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/challenge_response/" + url.PathEscape(request.StateKey) + "/" + url.PathEscape(request.PersonUID) + "/"

	req, err := http.NewRequest(http.MethodGet, requestURI, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req = req.WithContext(ctx)

	return req, err
}

// GetChallengeResponseStateKeyPersonUIDResponse is operation response value.
type GetChallengeResponseStateKeyPersonUIDResponse struct {
	StatusCode              int
	RawBody                 []byte                                                 // RawBody contains read bytes of response body.
	ValueOK                 *GetChallengeResponseStateKeyPersonUIDResponseValueOK  // ValueOK is a value of 200 OK response.
	ValueServiceUnavailable *APIErrors                                             // ValueServiceUnavailable is a value of 503 Service Unavailable response.
}

// decode loads data from *http.Response.
func (result *GetChallengeResponseStateKeyPersonUIDResponse) decode(resp *http.Response) error {
	var err error

	dump := bytes.NewBuffer(nil)
	body := io.TeeReader(resp.Body, dump)

	result.StatusCode = resp.StatusCode

	switch resp.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(body).Decode(&result.ValueOK)
	case http.StatusServiceUnavailable:
		err = json.NewDecoder(body).Decode(&result.ValueServiceUnavailable)
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

// GetChallengeResponseStateKeyPersonUID performs REST operation.
func (c *Client) GetChallengeResponseStateKeyPersonUID(ctx context.Context, request GetChallengeResponseStateKeyPersonUIDRequest) (result GetChallengeResponseStateKeyPersonUIDResponse, err error) {
	if c.InstrumentCtxFunc != nil {
		ctx = c.InstrumentCtxFunc(ctx, http.MethodGet, "/challenge_response/{state_key}/{person_uid}/", &request)
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