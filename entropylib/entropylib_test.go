package entropylib

import (
	"testing"
)

var EPSILON float64 = 0.000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

var CalculateShannonTests = []struct {
	in       string
	expected float64
}{
	{"", 0.0},
	{"ABCDEFGH", 3.0},
	{"AAAAAAAAAAAAAAAAAAAAAAAAAAAA", 0.0},
}

func TestCalculateShannon(t *testing.T) {
	for _, tt := range CalculateShannonTests {
		if actual := CalculateShannon(tt.in); actual != tt.expected {
			t.Errorf("CalculateShannon(%s): expected %f,actual %f", tt.in, tt.expected, actual)
		}
	}
}

var CalculateBalancedTests = []struct {
	in       string
	expected float64
}{
	{"", 0.0},
	{"ABCDEFGH", 1.0},
	{"AAAAAAAAAAAAAAAAAAAAAAAAAAAA", 0.0},
}

func TestCalculateBalanced(t *testing.T) {
	for _, tt := range CalculateBalancedTests {
		if actual := CalculateBalanced(tt.in); actual != tt.expected {
			t.Errorf("CalculateBalanced(%s): expected %f,actual %f", tt.in, tt.expected, actual)
		}
	}
}

var CalculateOnCharsetTests = []struct {
	in         string
	charsetLen int
	expected   float64
}{
	{"", 0, 0.0},
	{"ABCDEFGH", 8, 1.0},
	{"ABCDEFGH", 10, 0.903090},
	{"CDEFGH", 6, 1.0},
	{"AAAAAAAAAAAAAAAAAAAAAAAAAAAA", 4, 0.0},
}

func TestCalculateOnCharset(t *testing.T) {
	for _, tt := range CalculateOnCharsetTests {
		if actual := CalculateOnCharset(tt.in, tt.charsetLen); !floatEquals(actual, tt.expected) {
			t.Errorf("CalculateOnCharset(%s,%d): expected %f,actual %f", tt.in, tt.charsetLen, tt.expected, actual)
		}
	}
}

var CalculateOfFileTests = []struct {
	path         string
	min_expected float64
	max_expected float64
}{
	{"testRandom", 7.999, 8.0},
	{"testPNG", 7.9, 8.0},
	{"testText", 3.0, 5.0},
}

func TestCalculateOfFile(t *testing.T) {
	for _, tt := range CalculateOfFileTests {
		if actual, _ := CalculateOfFile(tt.path); actual > tt.max_expected || actual < tt.min_expected {
			t.Errorf("CalculateOfFile(%s): %f < actual < %f is FALSE", tt.path, tt.min_expected, actual, tt.max_expected)
		}
	}
}
