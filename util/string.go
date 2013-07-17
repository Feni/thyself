package util

import (
	"strconv"
)

// Limit a string to a max size
func Slice(str string, limit int) string {
	if len(str) > limit {
		return str[0:limit]
	}
	return str
}

// Convert a string to an int. If the string is not a valid int, return defaultVal
func ConvertToInt(rawVal string, defaultVal int64) int64 {
	if rawVal != "" {
		tempVal, err := strconv.ParseInt(rawVal, 10, 64) // base 10, 32 bit
		if err == nil {
			return tempVal
		}
	}
	return defaultVal
}
