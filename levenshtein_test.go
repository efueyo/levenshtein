package levenshtein_test

import (
	"testing"

	"github.com/efueyo/levenshtein"
)

func TestLDistnace(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		dist int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "b", 1},
		{"agricultor", "apicultor", 2},
		{"my dog is beautiful", "my cat is beautiful", 3},
	}
	for _, test := range tests {
		actual := levenshtein.LDistance(test.str1, test.str2)
		if actual != test.dist {
			t.Errorf("Distance(%q,%q) = %v; want %v", test.str1, test.str2, actual, test.dist)
		}
	}
}
