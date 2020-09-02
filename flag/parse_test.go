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
				Config: &Config{Port: 8080, NotFoundFile: ""},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "3000"},
			Expect: &Expect{
				Config: &Config{Port: 3000, NotFoundFile: ""},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "foo"},
			Expect: &Expect{
				Config: nil,
				Err:    errors.New(""),
			},
		}, {
			Args: []string{"-not-found-file"},
			Expect: &Expect{
				Config: nil,
				Err:    errors.New(""),
			},
		}, {
			Args: []string{"-not-found-file", "404.html"},
			Expect: &Expect{
				Config: &Config{Port: 8080, NotFoundFile: "404.html"},
				Err:    nil,
			},
		}, {
			Args: []string{"-not-found-file", "foo"},
			Expect: &Expect{
				Config: &Config{Port: 8080, NotFoundFile: "foo"},
				Err:    nil,
			},
		}, {
			Args: []string{"-port", "3000", "-not-found-file", "index.html"},
			Expect: &Expect{
				Config: &Config{Port: 3000, NotFoundFile: "index.html"},
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
		if got.Port != tc.Expect.Config.Port || got.NotFoundFile != tc.Expect.Config.NotFoundFile {
			t.Errorf(
				"For arguments %+v, expected %+v, but got %+v.",
				tc.Args,
				*tc.Expect.Config,
				*got,
			)
		}
	}
}
