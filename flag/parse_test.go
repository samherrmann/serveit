package flag

import (
	"errors"
	"testing"
)

type Tests []struct {
	Args   []string
	Expect *Expect
}

type Expect struct {
	Config *Config
	Err    error
}

func TestParse(t *testing.T) {
	// Define test cases
	tests := Tests{
		{
			Args: []string{},
			Expect: &Expect{
				Config: &Config{Port: 8080, SPAMode: false},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "3000"},
			Expect: &Expect{
				Config: &Config{Port: 3000, SPAMode: false},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "foo"},
			Expect: &Expect{
				Config: nil,
				Err:    errors.New(""),
			},
		}, {
			Args: []string{"-spa"},
			Expect: &Expect{
				Config: &Config{Port: 8080, SPAMode: true},
				Err:    nil,
			},
		}, {
			Args: []string{"-spa", "true"},
			Expect: &Expect{
				Config: &Config{Port: 8080, SPAMode: true},
				Err:    nil,
			},
		}, {
			Args: []string{"-spa", "foo"},
			Expect: &Expect{
				Config: &Config{Port: 8080, SPAMode: true},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "3000", "-spa"},
			Expect: &Expect{
				Config: &Config{Port: 3000, SPAMode: true},
				Err:    nil,
			},
		},
	}
	// Loop over all test cases
	for _, tc := range tests {
		got, err := parse(tc.Args)
		// Check if an error occured when no error is expected
		if err != nil && tc.Expect.Err == nil {
			t.Errorf("For arguments %+v, an error occured but no error was expected.", tc.Args)
			continue
		}
		// Check if no error occured when one is expected
		if err == nil && tc.Expect.Err != nil {
			t.Errorf("For arguments %+v, expected an error but no error occured.", tc.Args)
			continue
		}
		// If error occured when one is expected, then all is well. Exit before checking values.
		if err != nil && tc.Expect.Err != nil {
			continue
		}
		// Check values against Expected values
		if got.Port != tc.Expect.Config.Port || got.SPAMode != tc.Expect.Config.SPAMode {
			t.Errorf(
				"For arguments %+v, expected %+v, but got %+v.",
				tc.Args,
				*tc.Expect.Config,
				*got,
			)
		}
	}
}
