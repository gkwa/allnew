package main

import (
	"os"

	"github.com/taylormonacelli/{{ cookiecutter.project_slug }}"
)

func main() {
	code := {{ cookiecutter.project_slug }}.Execute()
	os.Exit(code)
}
