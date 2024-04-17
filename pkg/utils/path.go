package utils

import (
	"strconv"
	"strings"
)

func GetIdFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")
	return strconv.Atoi(parts[len(parts)-1])
}
