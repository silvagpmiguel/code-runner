package runner

import (
	"fmt"
	"os/exec"
	"strings"
)

// Compiler structure
type Compiler struct {
	Name       string
	Command    string
	InputPath  string
	OutputPath string
	To         string
	Extra      []string
}

// NewCompiler creates a new compiler for a specificic language
func NewCompiler(lang string, path string) (*Compiler, error) {
	compilers := map[string]string{
		"java":   "javac",
		"c":      "gcc",
		"kotlin": "kotlinc",
		"go":     "go",
	}
	commands := map[string]string{
		"java":   "",
		"c":      "",
		"kotlin": "-include-runtime",
		"go":     "build",
	}
	to := map[string]string{
		"java":   "-d",
		"c":      "-o",
		"kotlin": "-d",
		"go":     "-o",
	}

	if _, ok := compilers[lang]; !ok {
		return nil, fmt.Errorf("language not suported")
	}

	comp := Compiler{
		Name:       compilers[lang],
		Command:    commands[lang],
		InputPath:  path,
		OutputPath: "output/",
		To:         to[lang],
	}

	return &comp, nil
}

// Compile input and return the output path
func (comp *Compiler) Compile() (string, error) {
	args := buildCompilerArgs(comp)
	cmd := exec.Command(comp.Name, args...)
	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf("compiling error: %v %v %v", comp.Name, strings.Join(args, " "), err)
	}

	return comp.OutputPath, nil
}

func buildCompilerArgs(comp *Compiler) []string {
	var args []string

	switch comp.Name {
	case "go":
		comp.OutputPath += "main"
		args = append(args, comp.Command, comp.To, comp.OutputPath, comp.InputPath)
		args = append(args, comp.Extra...)
	case "kotlinc":
		comp.OutputPath += "main.jar"
		args = append(args, comp.InputPath, comp.Command, comp.To, comp.OutputPath)
		args = append(args, comp.Extra...)
	default:
		args = append(args, comp.InputPath, comp.To, comp.OutputPath)
		args = append(args, comp.Extra...)
	}

	return args
}
