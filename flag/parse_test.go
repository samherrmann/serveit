package flag_test

import (
	"errors"
	"testing"

	"github.com/samherrmann/serveit/flag"
)

func TestParse(t *testing.T) {
	type Want struct {
		config *flag.Config
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
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false},
				err:    nil,
			},
		}, {
			args: []string{"-port", "3000"},
			want: &Want{
				config: &flag.Config{Port: 3000, NotFoundFile: "", TLS: false},
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
				config: &flag.Config{Port: 8080, NotFoundFile: "404.html", TLS: false},
				err:    nil,
			},
		}, {
			args: []string{"-not-found-file", "foo"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "foo", TLS: false},
				err:    nil,
			},
		}, {
			args: []string{"-port", "3000", "-not-found-file", "index.html", "-tls"},
			want: &Want{
				config: &flag.Config{Port: 3000, NotFoundFile: "index.html", TLS: true},
				err:    nil,
			},
		}, {
			args: []string{"-tls"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: true},
				err:    nil,
			},
		}, {
			args: []string{"-tls=false"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false},
				err:    nil,
			},
		}, {
			args: []string{"-tls", "false"},
			want: &Want{
				// Note that setting boolean flag values explicitly needs to be done in
				// the form of "-flag=value", the form "-flag value" does not work.
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: true},
				err:    nil,
			},
		},
	}
	// Loop over all test cases
	for _, tc := range tests {
		got, err := flag.Parse(tc.args)
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
		// Check values against expected values
		if doValuesMatch(got, tc.want.config) {
			t.Errorf(
				"For arguments %+v, got %+v, but want %+v.",
				tc.args,
				*got,
				*tc.want.config,
			)
		}
	}
}

func doValuesMatch(got *flag.Config, want *flag.Config) bool {
	return got.Port != want.Port ||
		got.NotFoundFile != want.NotFoundFile ||
		got.TLS != want.TLS
}
