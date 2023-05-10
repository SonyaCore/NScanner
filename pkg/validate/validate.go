package validate

import (
	"errors"
)

var maxPort = 65535
var minPort = 0

func ValidateRange(num int) error {
	if num > maxPort || num <= minPort {
		return errors.New("Invalid port range")
	}
	return nil
}
