package validate

import (
	"errors"
	"fmt"
)

var protocols = []string{
	"udp", "tcp",
}

var maxPort = 65535
var minPort = 0

func ValRange(num int) error {
	if num > maxPort || num < minPort {
		return fmt.Errorf("Invalid port range %d", num)
	}
	return nil
}

func ValProtocols(protocol string) error {
	for _, proto := range protocols {
		if protocol == proto {
			return nil
		}
	}
	return errors.New("Invalid protocol")
}
