// utils/errors.go
package utils

import (
    "errors"
)

func NewError(message string) error {
    return errors.New(message)
}
