package services

import (
	"fmt"
	"testing"
)

func TestRatesService_Get(t *testing.T) {
	var tests = []struct {
		code          string
		rateExpected  float64
		errorExpected string
	}{
		{
			code:          "",
			errorExpected: "rate: code is empty",
			rateExpected:  0,
		},
		{
			code:          "XX",
			errorExpected: "rate: code not found",
			rateExpected:  0,
		},
		{
			code:          "ES",
			errorExpected: "",
			rateExpected:  21,
		},
	}

	ratesService := NewRatesService()

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			rate, err := ratesService.Get(test.code)
			if rate != nil && rate.Standard != test.rateExpected {
				t.Errorf("Test failed: wrong value\nwant: %b\ngot:  %b", test.rateExpected, rate.Standard)
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
