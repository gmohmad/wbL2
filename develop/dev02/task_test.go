package main

import (
	"errors"
	"testing"
)

type TestCase struct {
	input       string
	expectedStr string
	expectedErr error
}

func TestUnpackString(t *testing.T) {
	testCases := []TestCase{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
		{"qwe\\4\\5", "qwe45", nil},
		{"qwe\\45", "qwe44444", nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", nil},
	}

	for _, test := range testCases {
		res, err := unpackString(test.input)
		if res != test.expectedStr || !compareErrors(err, test.expectedErr) {
			t.Errorf("Test failed: expected - (%s, %v), got - (%s, %v)", test.expectedStr, test.expectedErr, res, err)
		} else {
			t.Log("Test Passed")
		}
	}
}

func compareErrors(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 != nil && err2 != nil {
		return err1.Error() == err2.Error()
	}
	return false
}
