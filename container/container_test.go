package container

import (
	"maps"
	"slices"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct{ got, want string }{
		{String(slices.Values([]int{1, 2, 3})), "[1 2 3]"},
		{String(slices.Values([]string{})), "[]"},
		{String2(maps.All(map[string]int{"a": 1})), "map[a:1]"},
		{String2(maps.All(map[string]int{})), "map[]"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %q, want %q", tt.got, tt.want)
		}
	}
}
