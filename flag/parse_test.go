package flag

import (
	"errors"
	"testing"
)

type Tests []struct {
	Args     []string
	Expected *Expected
}

type Expected struct {
	Config *Config
	Err    error
}

func TestParse(t *testing.T) {
	// Define test cases
	tests := Tests{
		{
			Args: []string{},
			Expected: &Expected{
				Config: &Config{Port: 8080, NotFoundFile: ""},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "3000"},
			Expected: &Expected{
				Config: &Config{Port: 3000, NotFoundFile: ""},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "foo"},
			Expected: &Expected{
				Config: nil,
				Err:    errors.New(""),
			},
		}, {
			Args: []string{"-not-found-file"},
			Expected: &Expected{
				Config: nil,
				Err:    errors.New(""),
			},
		}, {
			Args: []string{"-not-found-file", "404.html"},
			Expected: &Expected{
				Config: &Config{Port: 8080, NotFoundFile: "404.html"},
				Err:    nil,
			},
		}, {
			Args: []string{"-not-found-file", "foo"},
			Expected: &Expected{
				Config: &Config{Port: 8080, NotFoundFile: "foo"},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "3000", "-not-found-file", "index.html"},
			Expected: &Expected{
				Config: &Config{Port: 3000, NotFoundFile: "index.html"},
				Err:    nil,
			},
		},
	}
	// Loop over all test cases
	for _, tc := range tests {
		got, err := parse(tc.Args)
		// Check if an error occured when no error is expected
		if err != nil && tc.Expected.Err == nil {
			t.Errorf("For arguments %+v, an error occured but no error was expected.", tc.Args)
			continue
		}
		// Check if no error occured when one is expected
		if err == nil && tc.Expected.Err != nil {
			t.Errorf("For arguments %+v, expected an error but no error occured.", tc.Args)
			continue
		}
		// If error occured when one is expected, then all is well. Exit before checking values.
		if err != nil && tc.Expected.Err != nil {
			continue
		}
		// Check values against Expected values
		if got.Port != tc.Expected.Config.Port || got.NotFoundFile != tc.Expected.Config.NotFoundFile {
			t.Errorf(
				"For arguments %+v, expected %+v, but got %+v.",
				tc.Args,
				*tc.Expected.Config,
				*got,
			)
		}
	}
}
