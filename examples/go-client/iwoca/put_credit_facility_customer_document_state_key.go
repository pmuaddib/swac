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

// PutCreditFacilityCustomerDocumentStateKeyRequest is operation request value.
type PutCreditFacilityCustomerDocumentStateKeyRequest struct {
	Body     *PutCreditFacilityCustomerDocumentStateKeyRequestBody  // Body is a JSON request body.
	StateKey string                                                 // StateKey is a required `state_key` parameter in path.
}

// encode creates *http.Request for request data.
func (request *PutCreditFacilityCustomerDocumentStateKeyRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/credit_facility_customer_document/" + url.PathEscape(request.StateKey) + "/"

	body, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, requestURI, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req = req.WithContext(ctx)

	return req, err
}

// PutCreditFacilityCustomerDocumentStateKeyResponse is operation response value.
type PutCreditFacilityCustomerDocumentStateKeyResponse struct {
	StatusCode    int
	RawBody       []byte       // RawBody contains read bytes of response body.
	ValueAccepted interface{}  // ValueAccepted is a value of 202 Accepted response.
}

// decode loads data from *http.Response.
func (result *PutCreditFacilityCustomerDocumentStateKeyResponse) decode(resp *http.Response) error {
	var err error

	dump := bytes.NewBuffer(nil)
	body := io.TeeReader(resp.Body, dump)

	result.StatusCode = resp.StatusCode

	switch resp.StatusCode {
	case http.StatusAccepted:
		err = json.NewDecoder(body).Decode(&result.ValueAccepted)
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

// PutCreditFacilityCustomerDocumentStateKey performs REST operation.
func (c *Client) PutCreditFacilityCustomerDocumentStateKey(ctx context.Context, request PutCreditFacilityCustomerDocumentStateKeyRequest) (result PutCreditFacilityCustomerDocumentStateKeyResponse, err error) {
	if c.InstrumentCtxFunc != nil {
		ctx = c.InstrumentCtxFunc(ctx, http.MethodPut, "/credit_facility_customer_document/{state_key}/", &request)
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