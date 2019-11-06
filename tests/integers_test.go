package validator_test

import (
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
