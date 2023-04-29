package word

import "testing"

/**func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("heh") {
		t.Error(`IsPalindrome("heh") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
**/

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"hello", false},
		{"kayak", true},
		{"A man, a plan, a canal: Panama", true},
		{"heh", true},
	}

	for _, w := range tests {
		if got := IsPalindrome(w.input); got != w.want {
			t.Errorf("IsPalindrome(%q) is %v, should be %v", w.input, got, w.want)
		}
	}

}
