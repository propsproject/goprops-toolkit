package testutils

import "testing"

// TestCase helper test case wrapper struct, helps with ranging over an array of test cases
type TestCase struct {
	Name string
	Test func(t *testing.T)
}
