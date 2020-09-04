package flag

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	type Want struct {
		config *Config
		err    error
	}

	type Test struct {
		args []string
		want *Want
	}

	// Define test cases
	tests := []Test{
		{
			args: []string{},
			want: &Want{
				config: &Config{Port: 8080, NotFoundFile: ""},
				err:    nil,
			},
		}, {
			args: []string{"-port", "3000"},
			want: &Want{
				config: &Config{Port: 3000, NotFoundFile: ""},
				err:    nil,
			},
		}, {
			args: []string{"-port", "foo"},
			want: &Want{
				config: nil,
				err:    errors.New(""),
			},
		}, {
			args: []string{"-not-found-file"},
			want: &Want{
				config: nil,
				err:    errors.New(""),
			},
		}, {
			args: []string{"-not-found-file", "404.html"},
			want: &Want{
				config: &Config{Port: 8080, NotFoundFile: "404.html"},
				err:    nil,
			},
		}, {
			args: []string{"-not-found-file", "foo"},
			want: &Want{
				config: &Config{Port: 8080, NotFoundFile: "foo"},
				err:    nil,
			},
		}, {
			args: []string{"-port", "3000", "-not-found-file", "index.html"},
			want: &Want{
				config: &Config{Port: 3000, NotFoundFile: "index.html"},
				err:    nil,
			},
		},
	}
	// Loop over all test cases
	for _, tc := range tests {
		got, err := parse(tc.args)
		// Check if an error occured when no error is expected
		if err != nil && tc.want.err == nil {
			t.Errorf("For arguments %+v, want no error but got an error.", tc.args)
			continue
		}
		// Check if no error occured when one is expected
		if err == nil && tc.want.err != nil {
			t.Errorf("For arguments %+v, want an error but no error occured.", tc.args)
			continue
		}
		// If error occured when one is expected, then all is well. Exit before checking values.
		if err != nil && tc.want.err != nil {
			continue
		}
		// Check values against Expected values
		if got.Port != tc.want.config.Port || got.NotFoundFile != tc.want.config.NotFoundFile {
			t.Errorf(
				"For arguments %+v, got %+v, but want %+v.",
				tc.args,
				*got,
				*tc.want.config,
			)
		}
	}
}
