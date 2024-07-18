package stringuppercase

import "testing"

// TestToUpperCase tests the ToUpperCase function with different test cases.
func TestToUpperCase(t *testing.T) {
	// Defined the test cases taking input as string and expected output to be uppercase.
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},   // Basic lowercase to uppercase.
		{"pranay", "PRANAY"}, // Corrected expected result to match uppercase conversion.
		{"GoLang", "GOLANG"}, // Mixed case to uppercase.
		{"abc", "ABC"},       // Lowercase to uppercase.
		{"123", "123"},       // Numeric strings should remain unchanged.
	}
	// Iterating over each test case.
	for _, test := range tests {
		result := ToUpperCase(test.input) // Calling the ToUppecase function with the input from the given test cases.
		// Checks if the result is equal to the expected value.
		if result != test.expected {
			// If not, error is thrown with the input and expected value.
			t.Errorf("ToUpperCase(%q) = %q; expected: %q", test.input, result, test.expected)
		}
	}
}
