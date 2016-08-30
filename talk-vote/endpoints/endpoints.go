package endpoints

import (
	"errors"
)

var errUnableToAssertRequestType = errors.New("Unable to assert the requst to the correct request type")
var errInvalidVoteRequestValue = errors.New("invalid vote request value")

type Validater interface {
	Validate() error
}

func Validate(v Validater) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}
