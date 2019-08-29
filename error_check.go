package renum

import "strings"

// IsErrorUndefinedEnum is used to check if an error is because
// an enum value was undefined.
func IsErrorUndefinedEnum(err error) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), undefinedError) {
		return true
	}

	if strings.Contains(err.Error(), undefinedMessage) {
		return true
	}

	if val, ok := err.(Error); ok {
		if val.Code() == 0 {
			return true
		}
	}

	return false
}
