package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	tests := []struct {
		words    []string
		expected map[string][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			words: []string{"пятак", "Пятка", "тяпка", "листок", "слиток", "столик", "Тяпка"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "мама"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "мама", "амам"},
			expected: map[string][]string{
				"аамм":   {"амам", "мама"},
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, tt := range tests {
		got := FindAnagramSets(tt.words)

		// Compare maps
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("expected: %v, got: %v", tt.expected, got)
		}
	}
}
