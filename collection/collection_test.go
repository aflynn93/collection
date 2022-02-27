package collection

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	var actual int

	var objects = []string{"test alongword string 1", "$!@#s^g#$^ alongword", "another string !@#$!@% * @! ! alongword", "4sg alongword"}

	var tests = []struct {
		testCase string
		expected int
	}{
		{testCase: objects[0], expected: 0},
		{testCase: objects[1], expected: 1},
		{testCase: objects[2], expected: 2},
		{testCase: objects[3], expected: 3},
		{testCase: "a missing string", expected: -1},
		{testCase: "'", expected: -1},
		{testCase: "", expected: -1},
	}

	for _, test := range tests {
		actual = IndexOf(test.testCase, objects)

		assert.True(t, test.expected == actual, "Expected %d got %d for testcase [%s]", test.expected, actual, test.testCase)
	}
}

func TestContains(t *testing.T) {
	var actual bool

	var objects = []string{"test alongword string 1", "$!@#s^g#$^ alongword", "another string !@#$!@% * @! ! alongword", "4sg alongword"}

	var tests = []struct {
		testCase string
		expected bool
	}{
		{testCase: objects[0], expected: true},
		{testCase: objects[1], expected: true},
		{testCase: objects[2], expected: true},
		{testCase: objects[3], expected: true},
		{testCase: "a missing string", expected: false},
		{testCase: "'", expected: false},
		{testCase: "", expected: false},
	}

	for _, test := range tests {
		actual = Contains(test.testCase, objects)

		assert.True(t, test.expected == actual, "Expected %t got %t for testcase [%s]", test.expected, actual, test.testCase)
	}
}

func TestAny(t *testing.T) {
	var actual bool

	var objects = []string{"test alongword string 1", "$!@#s^g#$^ alongword", "another string !@#$!@% * @! ! alongword", "4sg alongword"}

	var tests = []struct {
		testCase string
		expected bool
	}{
		{testCase: "s", expected: true},
		{testCase: " alongword", expected: true},
		{testCase: "h", expected: true},
		{testCase: "@", expected: true},
		{testCase: "z", expected: false},
		{testCase: "p", expected: false},
		{testCase: "_", expected: false},
		{testCase: "missing", expected: false},
	}

	for _, test := range tests {
		actual = Any(objects, func(v string) bool {
			return strings.Contains(v, test.testCase)
		})

		assert.True(t, test.expected == actual, "strings.Contains test - Expected %t got %t for testcase [%s]", test.expected, actual, test.testCase)
	}
}

func TestAll(t *testing.T) {
	var actual bool

	var objects = []string{"test alongword string 1", "$!@#s^g#$^ alongword", "another string !@#$!@% * @! ! alongword", "4sg alongword"}

	var tests = []struct {
		testCase string
		expected bool
	}{
		{testCase: "s", expected: true},
		{testCase: " alongword", expected: true},
		{testCase: "h", expected: false},
		{testCase: "@", expected: false},
		{testCase: "z", expected: false},
		{testCase: "p", expected: false},
		{testCase: "_", expected: false},
		{testCase: "missing", expected: false},
	}

	for _, test := range tests {
		actual = All(objects, func(v string) bool {
			return strings.Contains(v, test.testCase)
		})

		assert.True(t, test.expected == actual, "strings.Contains test - Expected %t got %t for testcase [%s]", test.expected, actual, test.testCase)
	}
}

func TestFilter(t *testing.T) {
	var actual []string

	var objects = []string{"test alongword string 1", "$!@#s^g#$^ alongword", "another string !@#$!@% * @! ! alongword", "4sg alongword"}

	var tests = []struct {
		testCase string
		expected []string
	}{
		{testCase: "s", expected: objects},          // expecting all members
		{testCase: " alongword", expected: objects}, // expecting all members
		{testCase: "h", expected: objects[2:3]},     // expecting only the member with index 2
		{testCase: "@", expected: objects[1:3]},     // expecting only the members with index 1 and 2
		{testCase: "z", expected: []string{}},       // expecting an empty slice
		{testCase: "p", expected: []string{}},       // expecting an empty slice
		{testCase: "_", expected: []string{}},       // expecting an empty slice
		{testCase: "missing", expected: []string{}}, // expecting an empty slice
	}

	for _, test := range tests {
		actual = Filter(objects, func(v string) bool {
			return strings.Contains(v, test.testCase) // If the string v from objects contains the string test.testCase, return true
		})

		assert.True(t, reflect.DeepEqual(test.expected, actual), "strings.Contains test - Expected %v got %v for testcase [%s]", test.expected, actual, test.testCase)
	}
}

func TestMap(t *testing.T) {
	var actual []string

	var objects = []string{"a lowercase string", "AN UPPERCASE STRING", "a MiXeD cAsE sTrInG", "!@#^$%&*"}

	var tests = []struct {
		testCase func(string) string
		expected []string
	}{
		{testCase: strings.ToUpper, expected: []string{"A LOWERCASE STRING", "AN UPPERCASE STRING", "A MIXED CASE STRING", "!@#^$%&*"}},
		{testCase: strings.ToLower, expected: []string{"a lowercase string", "an uppercase string", "a mixed case string", "!@#^$%&*"}},
	}

	for _, test := range tests {
		actual = Map(objects, test.testCase)

		assert.True(t, reflect.DeepEqual(test.expected, actual), "strings.Contains test - Expected %v got %v for testcase [%s]", test.expected, actual, test.testCase)
	}
}
