package binary_operations

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		left   string
		right  string
		result string
	}{
		{
			"1101",
			"10",
			"1111",
		},
		{
			"1010",
			"1011",
			"10101",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.result, func(t *testing.T) {
			actual := AddBinary(testCase.left, testCase.right)
			if actual != testCase.result {
				t.Errorf("invalid result of addintion %s with %s", testCase.left, testCase.right)
			}
		})
	}

}
