package utils

import (
	"strconv"
)

func StringToUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
