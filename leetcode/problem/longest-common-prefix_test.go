package problem_test

import (
	"1d1c/leetcode/problem"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "prefix found",
			args: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "prefix not found",
			args: []string{"dog", "racecar", "car"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem.LongestCommonPrefix(tt.args); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
