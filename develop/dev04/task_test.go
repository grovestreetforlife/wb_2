package main

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	testTable := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:     "Test 1. Общий функционал",
			input:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{"пятак": {"пятка", "тяпка"}, "листок": {"слиток", "столик"}},
		},
		{
			name:     "Test 2. 2 слова из разных подмножеств",
			input:    []string{"тяпка", "листок"},
			expected: map[string][]string{},
		},
		{
			name:     "Test 3. Слова из одного подмножества с повтором",
			input:    []string{"тяпка", "пятка", "пятка"},
			expected: map[string][]string{},
		},
		{
			name:     "Test 4. 2 слова из одного подмножества",
			input:    []string{"тяпка", "пятка"},
			expected: map[string][]string{},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			actual := findAnagram(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("In %s, expected: %v, got: %v", tt.name, tt.expected, actual)
			}
		})
	}
}
