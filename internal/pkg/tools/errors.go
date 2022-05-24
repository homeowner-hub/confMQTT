package tools

import (
	"fmt"
	"math/rand"
)

type (
	ErrInvalidJSON struct {
		TypeName string
	}

	ErrKeyNotOptional struct {
		KeyName string
	}

	ErrDoubleErr struct {
		Err1 error
		Err2 error
	}
)

func (e *ErrInvalidJSON) Error() string {
	return fmt.Sprintf("json for '%s' is invalid", e.TypeName)
}

func (e *ErrKeyNotOptional) Error() string {
	return fmt.Sprintf("key '%s' is not optional", e.KeyName)
}

func (e *ErrDoubleErr) Error() string {
	return fmt.Sprintf("%s, %s", e.Err1.Error(), e.Err2.Error())
}

func (e *ErrDoubleErr) Unwrap() error {
	return []error{e.Err2, e.Err1}[rand.Intn(2)]
}
