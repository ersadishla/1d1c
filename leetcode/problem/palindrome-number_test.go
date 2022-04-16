package problem_test

import (
	"1d1c/leetcode/problem"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "121 is palindrome",
			args: 121,
			want: true,
		},
		{
			name: "-121 is not palindrome",
			args: -121,
			want: false,
		},
		{
			name: "10 is not palindrome",
			args: 10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem.IsPalindrome(tt.args); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
