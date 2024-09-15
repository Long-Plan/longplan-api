package errors

import "errors"

func Equal(err error, target error) bool {
	return errors.Is(err, target)
}
