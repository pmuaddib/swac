// Code generated by github.com/swaggest/swac v0.1.7, DO NOT EDIT.

package uspto

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ComponentsSchemasDataSetList structure is generated from "#/components/schemas/dataSetList".
type ComponentsSchemasDataSetList struct {
	Total                int64                                   `json:"total,omitempty"`
	Apis                 []ComponentsSchemasDataSetListApisItems `json:"apis,omitempty"`
	AdditionalProperties map[string]interface{}                  `json:"-"`               // All unmatched properties
}

type marshalComponentsSchemasDataSetList ComponentsSchemasDataSetList

var ignoreKeysComponentsSchemasDataSetList = []string{
	"total",
	"apis",
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSchemasDataSetList) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalComponentsSchemasDataSetList(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysComponentsSchemasDataSetList {
		delete(m, key)
	}

	for key, rawValue := range m {
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = ComponentsSchemasDataSetList(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSchemasDataSetList) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponentsSchemasDataSetList(i), i.AdditionalProperties)
}

// ComponentsSchemasDataSetListApisItems structure is generated from "#/components/schemas/dataSetList->apis->items".
type ComponentsSchemasDataSetListApisItems struct {
	APIKey               string                 `json:"apiKey,omitempty"`              // To be used as a dataset parameter value
	APIVersionNumber     string                 `json:"apiVersionNumber,omitempty"`    // To be used as a version parameter value
	APIURL               string                 `json:"apiUrl,omitempty"`              // The URL describing the dataset's fields
	APIDocumentationURL  string                 `json:"apiDocumentationUrl,omitempty"` // A URL to the API console for each API
	AdditionalProperties map[string]interface{} `json:"-"`                             // All unmatched properties
}

type marshalComponentsSchemasDataSetListApisItems ComponentsSchemasDataSetListApisItems

var ignoreKeysComponentsSchemasDataSetListApisItems = []string{
	"apiKey",
	"apiVersionNumber",
	"apiUrl",
	"apiDocumentationUrl",
}

// UnmarshalJSON decodes JSON.
func (i *ComponentsSchemasDataSetListApisItems) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalComponentsSchemasDataSetListApisItems(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysComponentsSchemasDataSetListApisItems {
		delete(m, key)
	}

	for key, rawValue := range m {
		if ii.AdditionalProperties == nil {
			ii.AdditionalProperties = make(map[string]interface{}, 1)
		}

		var val interface{}

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		ii.AdditionalProperties[key] = val
	}

	*i = ComponentsSchemasDataSetListApisItems(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i ComponentsSchemasDataSetListApisItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponentsSchemasDataSetListApisItems(i), i.AdditionalProperties)
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	isObject := true

	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		if string(j) == "{}" {
			continue
		}

		if string(j) == "null" {
			continue
		}

		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false

				continue
			}

			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}

		result = append(result, j[1:]...)
	}

	// Close empty result.
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
