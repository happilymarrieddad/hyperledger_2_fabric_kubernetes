// Package types defines all of the data entities this API works with.
// These may be used within repos, within API routes, etc...
package types

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	validCop *validator.Validate = validator.New()
)

func init() {
	validCop = validator.New()
}

// ValidatingEntity is an entity that can check itself for general validity errors
type ValidatingEntity interface {
	Validate() error
}

func Validate(entity interface{}) error {
	ve, ok := entity.(ValidatingEntity)
	if ok {
		return ve.Validate()
	}
	return validCop.Struct(entity)
}
