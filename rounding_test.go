package main

import (
	"testing"
)

// Test the BnmRoundStr function
// Test cases are acquired from https://www.bnm.gov.my/misc/-/asset_publisher/2BOPbOBfILtL/content/about-the-rounding-mechanism
func TestBnmRoundStr(t *testing.T) {
	testCases := map[string]string{
		"82.01": "82.00",
		"82.02": "82.00",
		"82.03": "82.05",
		"82.04": "82.05",
		"82.06": "82.05",
		"82.07": "82.05",
		"82.08": "82.10",
		"82.09": "82.10",
	}

	for amount, expected := range testCases {
		rounded, err := BnmRoundStr(amount)

		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if rounded != expected {
			t.Fatalf("expected %s to be rounded to %s, got %s\n", amount, expected, rounded)
		}
	}
}

func TestBnmRoundStrExtraTestCases(t *testing.T) {
	testCases := map[string]string{
		"82.11": "82.10",
	}

	for amount, expected := range testCases {
		rounded, err := BnmRoundStr(amount)

		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if rounded != expected {
			t.Fatalf("expected %s to be rounded to %s, got %s\n", amount, expected, rounded)
		}
	}
}

func TestBnmRoundStrInvalidAmount(t *testing.T) {
	invalidAmount1 := ""

	_, err := BnmRoundStr(invalidAmount1)

	if err == nil {
		t.Fatalf("expecting an error, but none was returned")
	}
}
