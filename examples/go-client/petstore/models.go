// Code generated by github.com/swaggest/swac v0.1.3, DO NOT EDIT.

package petstore

import (
	"bytes"
	"encoding/json"
	"errors"
)

// NewPet structure is generated from "#/definitions/NewPet".
type NewPet struct {
	Name                 string                 `json:"name,omitempty"`
	Tag                  string                 `json:"tag,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties
}

type marshalNewPet NewPet

var ignoreKeysNewPet = []string{
	"name",
	"tag",
}

// UnmarshalJSON decodes JSON.
func (i *NewPet) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalNewPet(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysNewPet {
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

	*i = NewPet(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i NewPet) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalNewPet(i), i.AdditionalProperties)
}

// PetAllOf1 structure is generated from "#/definitions/Pet/allOf/1".
type PetAllOf1 struct {
	ID                   int64                  `json:"id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`            // All unmatched properties
}

type marshalPetAllOf1 PetAllOf1

var ignoreKeysPetAllOf1 = []string{
	"id",
}

// UnmarshalJSON decodes JSON.
func (i *PetAllOf1) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalPetAllOf1(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysPetAllOf1 {
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

	*i = PetAllOf1(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i PetAllOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPetAllOf1(i), i.AdditionalProperties)
}

// Pet structure is generated from "#/definitions/Pet".
type Pet struct {
	NewPet *NewPet    `json:"-"`
	AllOf1 *PetAllOf1 `json:"-"`
}

// UnmarshalJSON decodes JSON.
func (i *Pet) UnmarshalJSON(data []byte) error {
	var err error

	err = json.Unmarshal(data, &i.NewPet)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &i.AllOf1)
	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON encodes JSON.
func (i Pet) MarshalJSON() ([]byte, error) {
	return marshalUnion(i.NewPet, i.AllOf1)
}

// Error structure is generated from "#/definitions/Error".
type Error struct {
	Code                 int64                  `json:"code,omitempty"`
	Message              string                 `json:"message,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                 // All unmatched properties
}

type marshalError Error

var ignoreKeysError = []string{
	"code",
	"message",
}

// UnmarshalJSON decodes JSON.
func (i *Error) UnmarshalJSON(data []byte) error {
	var err error

	ii := marshalError(*i)

	err = json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	var m map[string]json.RawMessage

	err = json.Unmarshal(data, &m)
	if err != nil {
		m = nil
	}

	for _, key := range ignoreKeysError {
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

	*i = Error(ii)

	return nil
}

// MarshalJSON encodes JSON.
func (i Error) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalError(i), i.AdditionalProperties)
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
