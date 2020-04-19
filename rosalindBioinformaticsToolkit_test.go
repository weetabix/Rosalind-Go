package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCountDNABases(t *testing.T) {
	tests := []struct {
		input      string
		wantResult map[string]int
	}{
		{"AGATCTT", map[string]int{
			"A": 2,
			"C": 1,
			"G": 1,
			"T": 3,
		},
		},
	}
	for _, tt := range tests {
		if gotResult := countDNABases(tt.input); !cmp.Equal(gotResult, tt.wantResult) {
			t.Errorf("countDNABases(%v) = %v, want %v", tt.input, gotResult, tt.wantResult)
		}
	}
}
