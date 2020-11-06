package algorithms

import "testing"

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		s      string
		t      string
		result bool
	}{
		{"nisheedh", "eedhnish", true},
		{"hello", "eedhnish", false},
	}

	for _, test := range tests {
		if result := isAnagram(test.s, test.t); result != test.result {
			t.Errorf("isAnagram(%s, %s) = %v", test.s, test.t, result)
		}
	}
}
