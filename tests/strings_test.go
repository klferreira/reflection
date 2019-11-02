package validator_test

import (
	"reflect"
	"testing"

	v "github.com/klferreira/valiant"
)

func TestStringValidationSuccess(t *testing.T) {
	schema := &v.Schema{
		"name": v.NewRules().String(),
	}

	user := map[string]interface{}{
		"name": "John",
	}

	ok, errs := v.Validate(user, schema)

	if !ok || len(errs) > 0 {
		t.Error("There were errors in the validation")
	}
}

func TestStringValidationError(t *testing.T) {
	ruleset := &v.Schema{
		"age": v.NewRules().String(),
	}

	user := map[string]interface{}{
		"age": 22,
	}

	ok, errs := v.Validate(user, ruleset)

	if ok || len(errs) == 0 {
		t.Error("There should be errors. None found.")
	}

	got := errs[0].Error()
	want := v.StringTypeErr("age", reflect.Int.String()).Error()

	if got != want {
		t.Error("The error message does not match the expected")
	}
}

func TestStringMaxLenValidationSuccess(t *testing.T) {
	ruleset := &v.Schema{
		"name": v.NewRules().String().MaxLen(4),
	}

	user := map[string]interface{}{
		"name": "John",
	}

	ok, errs := v.Validate(user, ruleset)
	if !ok || len(errs) > 0 {
		t.Error("Validation was expected to be successful, but errors were found")
	}
}

func TestStringMaxLenValidationError(t *testing.T) {
	ruleset := &v.Schema{
		"name": v.NewRules().String().MaxLen(3),
	}

	user := map[string]interface{}{
		"name": "John",
	}

	ok, errs := v.Validate(user, ruleset)
	if ok || errs == nil || len(errs) == 0 {
		t.Error("Validation was expected to have errors, but none were found")
	} else {
		got := errs[0].Error()
		want := v.StringMaxLenExceededErr("name", 3).Error()

		if got != want {
			t.Error("The error message does not match the expected")
		}
	}
}
