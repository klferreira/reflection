package validator

import "fmt"

func StringTypeErr(key string, vtype string) error {
	return fmt.Errorf("\"%s\" is of type %s, should be string", key, vtype)
}

func StringMaxLenExceededErr(key string, max int) error {
	return fmt.Errorf("\"%s\" exceede the maximum length of %d", key, max)
}
