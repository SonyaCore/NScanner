package util

import (
	"errors"
	"strconv"
	"strings"
)

func SplitPorts(listPort string) (int, int, error) {
	if !strings.Contains(listPort, "-") {
		return 0, 0, errors.New("range doesn't contain '-'")
	}
	rangePort := strings.Split(listPort, "-")

	startPort, _ := strconv.Atoi(rangePort[0])

	endPort, _ := strconv.Atoi(rangePort[1])

	return startPort, endPort, nil

}
