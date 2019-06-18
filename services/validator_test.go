package services

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		vatNumber     string
		foundExpected bool
		errorExpected string
	}{
		{
			vatNumber:     "",
			errorExpected: "vat: vat number is empty",
			foundExpected: false,
		},
		{
			vatNumber:     "ES1234123",
			errorExpected: "",
			foundExpected: false,
		},
	}

	validatorService := NewValidatorService()

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			found, err := validatorService.Validate(test.vatNumber)
			if found != test.foundExpected {
				t.Errorf("Test failed: wrong value\nwant: %t\ngot:  %t", test.foundExpected, found)
			}
			if err != nil && err.Error() != test.errorExpected {
				t.Errorf("Test failed: wrong error\nwant: %q\ngot:  %q", test.errorExpected, err.Error())
			}
			if err == nil && test.errorExpected != "" {
				t.Errorf("Test failed: no error\nexpected:  %q", test.errorExpected)
			}
		})
	}
}
