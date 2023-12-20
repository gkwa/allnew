package main

import (
	"os"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}"
)

func main() {
	code := {{ cookiecutter.project_slug }}.Execute()
	os.Exit(code)
}
