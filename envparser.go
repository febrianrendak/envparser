package envparser

import (
	"os"
	"strconv"
	"strings"
)

// GetEnvStr return value of env variable. Not null error returned if env variable is empty string
func GetEnvString(key string) (string, error) {
	isValid := ValidateKeyName(key)
	if isValid == false {
		return "", &ErrorWrapper{KeyName: key, Err: ErrInvalidVarName}
	}

	v := strings.TrimSpace(os.Getenv(key))
	if len(v) == 0 {
		return v, &ErrorWrapper{KeyName: key, Err: ErrEnvVarEmpty}
	}

	return v, nil
}

// GetEnvInt return value of env variable. Value will be set to 0 and error will be returned if env variable failed to parse
func GetEnvInt(key string) (int, error) {
	s, err := GetEnvString(key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return v, err
}

// GetEnvBool return value of env variable. Value will be set to false and error will be returned if env variable failed to parse
func GetEnvBool(key string) (bool, error) {
	s, err := GetEnvString(key)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}

	return v, err
}

// GetEnvFloat64 return value of env variable. Value will be set to 0.0 and error will be returned if env variable failed to parse
func GetEnvFloat64(key string) (float64, error) {
	s, err := GetEnvString(key)
	if err != nil {
		return 0.0, err
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}

	return v, err
}

// MustGetEnvStr return value of env variable if any or throw panic if blank or failed to parse variable
func MustGetEnvString(key string) string {
	v, err := GetEnvString(key)
	if err != nil {
		panic(err)
	}
	return v
}

// MustGetEnvInt return value of env variable if any or throw panic if blank or failed to parse variable
func MustGetEnvInt(key string) int {
	v, err := GetEnvInt(key)
	if err != nil {
		panic(err)
	}

	return v
}

// MustGetEnvBool return value of env variable if any or throw panic if blank or failed to parse variable
func MustGetEnvBool(key string) bool {
	v, err := GetEnvBool(key)
	if err != nil {
		panic(err)
	}

	return v
}

// MustGetEnvFloat64 return value of env variable if any or throw panic if blank or failed to parse variable
func MustGetEnvFloat64(key string) float64 {
	v, err := GetEnvFloat64(key)
	if err != nil {
		panic(err)
	}

	return v
}

// GetEnvStringWithDefault return value of env variable if any or return default value if error
func GetEnvStringWithDefault(key string, defaultValue string) string {
	v, err := GetEnvString(key)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetEnvIntWithDefault return value of env variable if any or return default value if error
func GetEnvIntWithDefault(key string, defaultValue int) int {
	v, err := GetEnvInt(key)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetEnvBoolWithDefault return value of env variable if any or return default value if error
func GetEnvBoolWithDefault(key string, defaultValue bool) bool {
	v, err := GetEnvBool(key)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetEnvFloat64WithDefault return value of env variable if any or return default value if error
func GetEnvFloat64WithDefault(key string, defaultValue float64) float64 {
	v, err := GetEnvFloat64(key)
	if err != nil {
		return defaultValue
	}
	return v
}
