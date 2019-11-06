package validator_test

import (
	"reflect"
	"testing"

	v "github.com/klferreira/valiant"
)

func TestIntValidationSuccess(t *testing.T) {
	schema := &v.Schema{
		"age": v.NewRules().Int(),
	}

	user := map[string]interface{}{
		"age": 22,
	}

	ok, errs := v.Validate(user, schema)

	if !ok || len(errs) > 0 {
		t.Error("There were errors in the validation")
	}
}

func TestIntValidationError(t *testing.T) {
	schema := &v.Schema{
		"age": v.NewRules().Int(),
	}

	user := map[string]interface{}{
		"age": "22",
	}

	ok, errs := v.Validate(user, schema)

	if ok || len(errs) == 0 {
		t.Error("There should be errors. None found.")
	}

	got := errs[0].Error()
	want := v.IntTypeErr("age", reflect.String.String()).Error()

	if got != want {
		t.Error("The error message does not match the expected")
	}
}

func TestIntMaxValidationSuccess(t *testing.T) {
	schema := &v.Schema{
		"age": v.NewRules().Int().Max(22),
	}

	user := map[string]interface{}{
		"age": 22,
	}

	ok, errs := v.Validate(user, schema)

	if !ok || len(errs) > 0 {
		t.Error("Validation was expected to be successful, but errors were found")
	}
}

func TestIntMaxValidationError(t *testing.T) {
	schema := &v.Schema{
		"age": v.NewRules().Int().Max(22),
	}

	user := map[string]interface{}{
		"age": 23,
	}

	ok, errs := v.Validate(user, schema)
	if ok || errs == nil || len(errs) == 0 {
		t.Error("Validation was expected to have errors, but none were found")
	} else {
		got := errs[0].Error()
		want := v.IntMaxLengthErr("age", 22).Error()

		if got != want {
			t.Error("The error message does not match the expected")
		}
	}
}
