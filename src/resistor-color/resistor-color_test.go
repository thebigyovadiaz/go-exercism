package resistor_color

import "testing"

var colorCodeTestCases = []struct {
	description string
	input       string
	expected    int
}{
	{"black", "black", 0},
	{"brown", "brown", 1},
	{"red", "red", 2},
	{"orange", "orange", 3},
	{"yellow", "yellow", 4},
	{"green", "green", 5},
	{"blue", "blue", 6},
	{"violet", "violet", 7},
	{"grey", "grey", 8},
	{"white", "white", 9},
}

var colorsTestCases = []struct {
	description string
	expected    []string
}{
	{"all colors", []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}},
}

func TestColorCode(t *testing.T) {
	for _, tc := range colorCodeTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := ColorCode(tc.input)
			if actual != tc.expected {
				t.Fatalf("ColorCode(%q): expected %d, actual %d", tc.input, tc.expected, actual)
			}
		})
	}
}

func TestColors(t *testing.T) {
	for _, tc := range colorsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Colors()
			if len(actual) != len(tc.expected) {
				t.Fatalf("Colors(): expected %+v, actual %+v", tc.expected, actual)
			}
			expectedMap := makeMap(tc.expected)
			actualMap := makeMap(actual)
			if !mapsEqual(expectedMap, actualMap) {
				t.Fatalf("Colors(): expected %+v, actual %+v", tc.expected, actual)
			}
		})
	}
}

// colorCodeBench is intended to be used in BenchmarkColorCode to avoid compiler optimizations.
var colorCodeBench int

func BenchmarkColorCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range colorCodeTestCases {
			colorCodeBench = ColorCode(tc.input)
		}
	}
}

// colorsBench is intended to be used in BenchmarkColors to avoid compiler optimizations.
var colorsBench []string

func BenchmarkColors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colorsBench = Colors()
	}
}

func makeMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}

func mapsEqual(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}
