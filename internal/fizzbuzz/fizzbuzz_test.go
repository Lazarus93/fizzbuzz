package fizzbuzz

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestGenerateSequence(t *testing.T) {
	service := NewService()

	tests := []struct {
		name           string
		string1        string
		string2        string
		int1           int
		int2           int
		limit          int
		expectedResult []string
		expectedError  error
	}{
		{
			name:    "Test 1",
			string1: "fizz",
			string2: "buzz",
			int1:    3,
			int2:    5,
			limit:   15,
			expectedResult: []string{
				"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz",
				"11", "fizz", "13", "14", "fizzbuzz",
			},
		},
		{
			name:    "Test 2",
			string1: "foo",
			string2: "bar",
			int1:    2,
			int2:    3,
			limit:   6,
			expectedResult: []string{
				"1", "foo", "bar", "foo", "5", "foobar",
			},
		},
		{
			name:           "Negative limit",
			string1:        "fizz",
			string2:        "buzz",
			int1:           3,
			int2:           5,
			limit:          -1,
			expectedResult: nil,
			expectedError:  errors.New("limit must be positive"),
		},
		{
			name:           "Zero int1",
			string1:        "fizz",
			string2:        "buzz",
			int1:           0,
			int2:           5,
			limit:          10,
			expectedResult: nil,
			expectedError:  errors.New("int1 and int2 must be positive"),
		},
		{
			name:           "Zero int2",
			string1:        "fizz",
			string2:        "buzz",
			int1:           2,
			int2:           0,
			limit:          10,
			expectedResult: nil,
			expectedError:  errors.New("int1 and int2 must be positive"),
		},
		{
			name:           "Empty string1",
			string1:        "",
			string2:        "buzz",
			int1:           3,
			int2:           5,
			limit:          15,
			expectedResult: nil,
			expectedError:  errors.New("string1 and string2 must be non-empty"),
		},
		{
			name:           "Too long string2",
			string1:        "fizz",
			string2:        strings.Repeat("a", 1001),
			int1:           3,
			int2:           5,
			limit:          15,
			expectedResult: nil,
			expectedError:  errors.New("string1 and string2 must be less than 1000 characters"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GenerateSequence(tt.string1, tt.string2, tt.int1, tt.int2, tt.limit)
			if err != nil && tt.expectedError == nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if err == nil && tt.expectedError != nil {
				t.Errorf("Expected error, but got none")
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError, err)
			}
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
