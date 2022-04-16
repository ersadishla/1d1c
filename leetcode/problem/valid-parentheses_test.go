package problem_test

import (
	"1d1c/leetcode/problem"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "valid parentheses",
			args: "({})",
			want: true,
		},
		{
			name: "invalid parentheses",
			args: "[(){]}",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem.IsValidParentheses(tt.args); got != tt.want {
				t.Errorf("IsValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
