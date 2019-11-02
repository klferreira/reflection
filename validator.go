package validator

import (
	"errors"
	"reflect"
)

type ValidationSchema interface {
	ValidateMap(v reflect.Value) (ok bool, errs []error)
}

type Schema map[string]RuleSet

func (s *Schema) GetFieldRules(key string) RuleSet {
	schema := (*s)
	return schema[key]
}

func (s *Schema) ValidateMap(v reflect.Value) (ok bool, errs []error) {
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key).Interface()
		keyErrs := s.GetFieldRules(key.String()).Validate(key.String(), value)
		if len(keyErrs) > 0 {
			errs = append(errs, keyErrs...)
		}
	}

	return len(errs) == 0, errs
}

func Validate(obj interface{}, sc ValidationSchema) (bool, []error) {
	switch v := reflect.ValueOf(obj); v.Kind() {
	case reflect.Map:
		return sc.ValidateMap(v)
	default:
		return false, []error{
			errors.New("Object is not of type map[string]interface{}"),
		}
	}
}
