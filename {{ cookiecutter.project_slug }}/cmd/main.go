package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
		os.Exit(0)
	}

	versionFlag := flag.Bool("version", false, "Print version information")
	flag.BoolVar(versionFlag, "v", false, "Print version information")

	flag.Parse()

	if *versionFlag {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
		os.Exit(0)
	}

	code := {{ cookiecutter.project_slug }}.Execute()
	os.Exit(code)
}
