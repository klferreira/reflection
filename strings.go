package validator

import (
	"reflect"
)

type StringRuleSet struct {
	maxLen int
	minLen int
}

func (r *StringRuleSet) GetType() string {
	return reflect.String.String()
}

func (r *StringRuleSet) checkType(key string, value reflect.Value) error {
	vtype := value.Kind().String()
	if vtype != reflect.String.String() {
		return StringTypeErr(key, vtype)
	}

	return nil
}

func (r *StringRuleSet) MaxLen(n int) *StringRuleSet {
	r.maxLen = n
	return r
}

func (r *StringRuleSet) checkMaxLen(key string, value string) error {
	if r.maxLen != 0 && len(value) > r.maxLen {
		return StringMaxLengthErr(key, r.maxLen)
	}

	return nil
}

func (r *StringRuleSet) MinLen(n int) *StringRuleSet {
	r.minLen = n
	return r
}

func (r *StringRuleSet) checkMinLen(key string, value string) error {
	if r.minLen != 0 && len(value) < r.minLen {
		return StringMinLengthErr(key, r.minLen)
	}

	return nil
}

func (r *StringRuleSet) Validate(key string, value interface{}) (errs []error) {
	rValue := reflect.ValueOf(value)
	sValue := rValue.String()

	if err := r.checkType(key, rValue); err != nil {
		errs = append(errs, err)
	}

	if err := r.checkMaxLen(key, sValue); err != nil {
		errs = append(errs, err)
	}

	if err := r.checkMinLen(key, sValue); err != nil {
		errs = append(errs, err)
	}

	return errs
}
