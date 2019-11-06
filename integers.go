package validator

import "reflect"

type IntRuleSet struct {
	max int64
}

func (i *IntRuleSet) checkType(key string, value reflect.Value) error {
	vtype := value.Kind().String()
	if vtype != reflect.Int.String() {
		return IntTypeErr(key, vtype)
	}

	return nil
}

func (i *IntRuleSet) Max(m int64) *IntRuleSet {
	i.max = m
	return i
}

func (i *IntRuleSet) checkMax(key string, iValue int64) error {
	if i.max != 0 && iValue > i.max {
		return IntMaxLengthErr(key, i.max)
	}

	return nil
}

func (i *IntRuleSet) Validate(key string, value interface{}) (errs []error) {
	rValue := reflect.ValueOf(value)

	if err := i.checkType(key, rValue); err != nil {
		errs = append(errs, err)
		return
	}

	iValue := rValue.Int()

	if err := i.checkMax(key, iValue); err != nil {
		errs = append(errs, err)
	}

	return errs
}
