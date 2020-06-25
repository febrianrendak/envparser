package envparser

import (
	"os"
	"testing"
)

func TestGetEnvString(t *testing.T) {
	var key = "ENVPARSER_TEST_KEY"
	var val = "ANY"
	os.Setenv(key, val)

	v, err := GetEnvString(key)
	if err != nil {
		t.Errorf("Result was incorrect, error should be nil but instead got error: %s", err)
	}

	if v != val {
		t.Errorf(`Result was incorrect, env var should be "%s" but instead got: "%s"`, val, v)
	}

	os.Unsetenv(key)
	v1, err1 := GetEnvString(key)
	if err1 == nil {
		t.Errorf("Result was incorrect, error should be not nil but instead got error: %v", nil)
	}

	if v1 != "" {
		t.Errorf(`Result was incorrect, env var should be empty string but instead got: "%s"`, v1)
	}
}

func TestEnvName(t *testing.T) {
	var keyName = `"must be invalid`
	isValid1 := ValidateKeyName(keyName)

	if isValid1 == true {
		t.Errorf(`Result was incorrect, env name "%s" should be invalid but the result stated as valid`, keyName)
	}

	keyName = "MUST_BE_VALID"
	isValid2 := ValidateKeyName(keyName)
	if isValid2 == false {
		t.Errorf(`Result was incorrect, env name "%s" should be valid but the result stated as invalid`, keyName)
	}
}
