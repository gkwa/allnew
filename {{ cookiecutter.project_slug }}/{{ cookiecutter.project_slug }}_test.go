package {{ cookiecutter.project_slug }}_test

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"github.com/taylormonacelli/{{ cookiecutter.project_slug }}"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"{{ cookiecutter.project_slug }}": {{ cookiecutter.project_slug }}.Execute,
	}))
}

func TestHello(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
