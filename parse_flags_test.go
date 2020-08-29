package main

import (
	"testing"
)

func TestParseFlags(t *testing.T) {
	// Define test cases
	tests := []struct {
		args   []string
		expect *Config
	}{
		{
			args:   []string{},
			expect: &Config{Port: 8080, SPAMode: false},
		},
		{
			args:   []string{"-port", "3000"},
			expect: &Config{Port: 3000, SPAMode: false},
		},
		{
			args:   []string{"-spa"},
			expect: &Config{Port: 8080, SPAMode: true},
		}, {
			args:   []string{"-port", "3000", "-spa"},
			expect: &Config{Port: 3000, SPAMode: true},
		},
	}
	// Loop over all test cases
	for _, tc := range tests {
		got := parseFlags(tc.args)
		// Check values against expected values
		if got.Port != tc.expect.Port || got.SPAMode != tc.expect.SPAMode {
			t.Errorf(
				"For arguments %v, expected %v, but got %v",
				tc.args,
				*tc.expect,
				*got,
			)
		}
	}
}
