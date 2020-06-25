package envparser

import (
	"regexp"
)

var pattern = "^[a-zA-Z_]+[a-zA-Z0-9_]*$"
var IsLetter = regexp.MustCompile(pattern).MatchString

// ValidateKeyName must ensure env variable name only contain alphanumeric + underscore
// return true if env variable name is valid otherwise return false
func ValidateKeyName(key string) bool {
	result := IsLetter(key)
	return result
}
