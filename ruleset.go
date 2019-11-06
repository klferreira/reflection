package validator

import (
	"fmt"
)

type RuleSet interface {
	Validate(key string, value interface{}) (errs []error)
}

type DefaultRuleSet struct{}

func NewRules() *DefaultRuleSet {
	return &DefaultRuleSet{}
}

func (r *DefaultRuleSet) String() *StringRuleSet {
	return &StringRuleSet{}
}

func (r *DefaultRuleSet) Int() *IntRuleSet {
	return &IntRuleSet{}
}

func (r *DefaultRuleSet) Validate(key string, value interface{}) (errs []error) {
	errs = append(errs, fmt.Errorf("Empty RuleSet for %s", key))
	return
}
