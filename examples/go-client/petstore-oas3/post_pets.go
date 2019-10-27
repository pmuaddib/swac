// Code generated by github.com/swaggest/swac v0.1.2, DO NOT EDIT.

package petstore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// PostPetsRequest is operation request value.
type PostPetsRequest struct {
	Body *ComponentsSchemasNewPet  // Body is a JSON request body.
}

// encode creates *http.Request for request data.
func (request *PostPetsRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestURI := baseURL + "/pets"

	body, err := json.Marshal(request.Body)
	if err != nil {
	    return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, requestURI, bytes.NewBuffer(body))
	if err != nil {
	    return nil, err
	}

	req = req.WithContext(ctx)

	return req, err
}

// PostPetsResponse is operation response value.
type PostPetsResponse struct {
	StatusCode int
	ValueOK    *ComponentsSchemasPet  // ValueOK is a value of 200 OK response.
}

// decode loads data from *http.Response.
func (result *PostPetsResponse) decode(resp *http.Response) error {
	result.StatusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK:
	    err := json.NewDecoder(resp.Body).Decode(&result.ValueOK)
	    if err != nil {
	        return err
	    }
	default:
	    return errors.New("unexpected response status: " + resp.Status)
	}

	return nil
}

// PostPets performs REST operation.
func (c *Client) PostPets(ctx context.Context, request PostPetsRequest) (PostPetsResponse, error) {
	result := PostPetsResponse{}
	ctx = context.WithValue(ctx, "restOperationPath", "/pets")
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
	defer resp.Body.Close()
	err = result.decode(resp)
	if err != nil {
	    return result, err
	}
	return result, nil
}