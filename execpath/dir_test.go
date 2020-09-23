package execpath_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/samherrmann/serveit/execpath"
)

func TestDir(t *testing.T) {
	exe, err := os.Executable()
	if err != nil {
		t.Error(err)
	}
	want := filepath.Dir(filepath.ToSlash(exe))

	got, err := execpath.Dir()
	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Wrong executable directory: Got %v, want %v", got, want)
	}
}
