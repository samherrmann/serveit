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
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-port", "3000"},
			want: &Want{
				config: &flag.Config{Port: 3000, NotFoundFile: "", TLS: false, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-port", "foo"},
			want: &Want{
				config: nil,
				err:    errors.New(""),
			},
		}, {
			args: []string{"program-name", "-not-found-file"},
			want: &Want{
				config: nil,
				err:    errors.New(""),
			},
		}, {
			args: []string{"program-name", "-not-found-file", "404.html"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "404.html", TLS: false, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-not-found-file", "foo"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "foo", TLS: false, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-port", "3000", "-not-found-file", "index.html", "-tls"},
			want: &Want{
				config: &flag.Config{Port: 3000, NotFoundFile: "index.html", TLS: true, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-tls"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: true, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-tls=false"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-tls", "false"},
			want: &Want{
				// Note that setting boolean flag values explicitly needs to be done in
				// the form of "-flag=value", the form "-flag value" does not work.
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: true, Hosts: []string{"localhost"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-hosts", "example.com"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false, Hosts: []string{"example.com"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-hosts", "localhost,example.com"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false, Hosts: []string{"localhost", "example.com"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-hosts", "localhost,example.com,192.168.0.1"},
			want: &Want{
				config: &flag.Config{Port: 8080, NotFoundFile: "", TLS: false, Hosts: []string{"localhost", "example.com", "192.168.0.1"}},
				err:    nil,
			},
		}, {
			args: []string{"program-name", "-hosts"},
			want: &Want{
				config: nil,
				err:    errors.New(""),
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
		got.TLS != want.TLS ||
		!equal(got.Hosts, want.Hosts)
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
