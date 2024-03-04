package common

import (
	"github.com/frank269/protoc-gen-code/proto/validate"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

// Disabled returns true if validations are disabled for msg
func Disabled(msg pgs.Message) (disabled bool, err error) {
	_, err = msg.Extension(validate.E_Disabled, &disabled)
	return
}

// RequiredOneOf returns true if the oneof field requires a field to be set
func RequiredOneOf(oo pgs.OneOf) (required bool, err error) {
	_, err = oo.Extension(validate.E_Required, &required)
	return
}
