package gotility

import (
	"regexp"
)

const (
	ASCII_VALUE_0 = 48
	ASCII_VALUE_9 = 57
	EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

func IsNumber (str string) bool {
	if len(str) == 0 {
		return false
	}
	for index := range str {
		if uint8(str[index]) < ASCII_VALUE_0 || uint8(str[index]) > ASCII_VALUE_9 {
			return false
		}
	}
	return true
}

func IsEmail (str string) bool {
	if len(str) == 0 {
		return false
	}
	var validEmail = regexp.MustCompile(EMAIL_REGEX)
	return validEmail.MatchString(str)
}




