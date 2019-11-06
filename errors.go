package validator

import "fmt"

func StringTypeErr(key string, vtype string) error {
	return fmt.Errorf("\"%s\" is of type %s, should be string", key, vtype)
}

func StringMaxLengthErr(key string, max int) error {
	return fmt.Errorf("\"%s\" length exceeded the specified maximum of %d", key, max)
}

func StringMinLengthErr(key string, min int) error {
	return fmt.Errorf("\"%s\" length is less than specified minimum of %d", key, min)
}

func IntTypeErr(key string, vtype string) error {
	return fmt.Errorf("\"%s\" is of type %s, should be int", key, vtype)
}

func IntMaxLengthErr(key string, max int64) error {
	return fmt.Errorf("\"%s\" length exceeded the specified maximum of %d", key, max)
}
