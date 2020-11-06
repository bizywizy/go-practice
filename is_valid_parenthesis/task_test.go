package is_valid_parenthesis

import (
	"testing"
)

func assert(t *testing.T, fnc func(string) bool, expected bool, value string) {
	if fnc(value) != expected {
		t.Error("Value is not true " + value)
	}
}

func TestIsValidParenthesis(t *testing.T) {

	assert(t, IsValid, true, "()")
	assert(t, IsValid, true, "()[]{}")
	assert(t, IsValid, false, "(]")
	assert(t, IsValid, false, "([)]")
	assert(t, IsValid, true, "{[]}")
	assert(t, IsValid, true, "(([]){})")
}
