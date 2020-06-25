package envparser

import (
	"testing"
)

func TestError(t *testing.T) {
	var keyName = "test test"

	_, err := GetEnvString(keyName)

	mustBeFalse := ErrorIsEnvVarEmpty(err)
	if mustBeFalse == true {
		t.Errorf(`error type is not matching. error should be "%s" but result is "%s"`, ErrInvalidVarName, err)
	}

	mustBeTrue := ErrorIsInvalidVarName(err)
	if mustBeTrue == false {
		t.Errorf(`error type is not matching. error should be "%s" but result is "%s"`, ErrInvalidVarName, err)
	}
}
