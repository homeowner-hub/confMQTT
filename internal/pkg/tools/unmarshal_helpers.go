package tools

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	JSONValue  []byte
	JSONObject map[string]JSONValue
)

func (j *JSONValue) UnmarshalJSON(b []byte) error {
	*j = b
	return nil
}

func UnmarshalBiKeys(typename string, b []byte, keys ...Key[any, any]) error {
	j, err := getJSON(b, typename)
	if err != nil {
		return err
	}

	for _, k := range keys {
		if err := UnmarshalBiKey(j, k); err != nil {
			return err
		}
	}
	return nil
}

func getJSON(b []byte, typename string) (JSONObject, error) {
	var j JSONObject
	if err := json.Unmarshal(b, &j); err != nil {
		return nil, &ErrInvalidJSON{
			TypeName: typename,
		}
	}
	return j, nil
}

func UnmarshalBiKey(obj JSONObject, key Key[any, any]) error {
	if raw, ok := getRaw(obj, key); ok {
		return json.Unmarshal(raw, &key)
	}
	return checkOptional(key)
}

func checkOptional(key Key[any, any]) error {
	if !key.Optional() {
		return &ErrKeyNotOptional{
			KeyName: key.LongName(),
		}
	}
	return nil
}

func getRaw(obj JSONObject, key Key[any, any]) ([]byte, bool) {
	if v, ok := obj[key.ShortName()]; ok {
		return v, true
	} else if v, ok := obj[key.LongName()]; ok {
		return v, true
	}
	return nil, false
}

type ValueUnmarshaler[T any] struct {
	Typename string
	Optional bool
	Value1   *T
	Exists1  *bool
}

type BiValueUnmarshaler[T, V any] struct {
	ValueUnmarshaler[T]
	Value2  *V
	Exists2 *bool
}

func (vu ValueUnmarshaler[T]) UnmarshalJSON(b []byte) error {
	if err := validJSON(vu.Typename, b); err != nil {
		return err
	}
	return vu.unmarshal(b)
}

func (vu ValueUnmarshaler[T]) unmarshal(b []byte) error {
	if err := json.Unmarshal(b, &vu.Value1); err != nil {
		if match := new(json.UnmarshalTypeError); errors.As(err, &match) {
			v, ok := interface{}(vu.Value1).(*string)
			if ok {
				var i interface{}
				if err := json.Unmarshal(b, &i); err != nil {
					return err
				}

				*v = fmt.Sprintf("%+v", i)
				*vu.Exists1 = true
				return nil
			}
		}

		return fmt.Errorf("can't unmarshal key '%s', %w", vu.Typename, err)
	}
	*vu.Exists1 = true
	return nil
}

func (vu BiValueUnmarshaler[T, V]) UnmarshalJSON(b []byte) error {
	if err := validJSON(vu.Typename, b); err != nil {
		return err
	}

	err1 := ValueUnmarshaler[T]{
		Typename: vu.Typename,
		Optional: vu.Optional,
		Value1:   vu.Value1,
		Exists1:  vu.Exists1,
	}.UnmarshalJSON(b)

	err2 := ValueUnmarshaler[V]{
		Typename: vu.Typename,
		Optional: vu.Optional,
		Value1:   vu.Value2,
		Exists1:  vu.Exists2,
	}.UnmarshalJSON(b)

	if err1 != nil && err2 != nil {
		return &ErrDoubleErr{
			Err1: err1,
			Err2: err2,
		}
	}

	if *vu.Exists2 {
		*vu.Exists1 = false
		setToZeroVal[T](vu.Value1)
	}

	return nil
}

func setToZeroVal[T any](p *T) {
	*p = *new(T)
}

func validJSON(typename string, b []byte) error {
	if json.Unmarshal(b, new(any)) != nil {
		return &ErrInvalidJSON{
			TypeName: typename,
		}
	}
	return nil
}
