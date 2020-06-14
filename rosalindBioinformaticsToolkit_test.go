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

func TestTransDNAToRNA(t *testing.T) {
	tests := []struct {
		input      string
		wantResult string
	}{
		{"GATGGAACTTGACTACGTAAATT", "GAUGGAACUUGACUACGUAAAUU"},
	}
	for _, tt := range tests {
		if gotResult := transDNAToRNA(tt.input); !cmp.Equal(gotResult, tt.wantResult) {
			t.Errorf("countDNABases(%v) = %v, want %v", tt.input, gotResult, tt.wantResult)
		}
	}
}

//func TestGCContent(t *testing.T) { //TODO fix test, set up yml source for test data.
//	tests := []struct {
//		input      string
//		wantResult float64
//	}{
//		{">Foo\nAGTCA\nTGCT", float64(44.444444)},
//		{"AGTCATGCC", float64(0)},
//	}
//	for _, tt := range tests {
//		if gotResult := gcContent(tt.input); !cmp.Equal(gotResult, tt.wantResult) {
//			t.Errorf("gcContent(%v) = %v, want %v", tt.input, gotResult, tt.wantResult)
//		}
//	}
//}

//AAAACCCGGT
//ACCGGGTTTT
