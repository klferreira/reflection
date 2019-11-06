package validator

import "reflect"

type IntRuleSet struct {
}

func (i *IntRuleSet) checkType(key string, value reflect.Value) error {
	vtype := value.Kind().String()
	if vtype != reflect.Int.String() {
		return IntTypeErr(key, vtype)
	}

	return nil
}

func (i *IntRuleSet) Validate(key string, value interface{}) (errs []error) {
	rValue := reflect.ValueOf(value)
	// iValue := rValue.Int()

	if err := i.checkType(key, rValue); err != nil {
		errs = append(errs, err)
	}

	return errs
}
