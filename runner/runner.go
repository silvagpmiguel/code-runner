package runner

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Runner structure
type Runner struct {
	Name      string
	Language  string
	Command   string
	InputPath string
	Extra     []string
}

// NewRunner creates a new code runner
func NewRunner(language string, inputPath string) (*Runner, error) {
	lang := strings.ToLower(language)

	if lang != "python" {
		comp, err := NewCompiler(lang, inputPath)

		if err != nil {
			return nil, fmt.Errorf("creating compiler: %v", err)
		}

		inputPath, err = comp.Compile()

		if err != nil {
			return nil, err
		}
	}

	runners := map[string]string{
		"java":   "java",
		"go":     "/bin/bash",
		"c":      "/bin/bash",
		"kotlin": "java",
		"python": "python",
	}
	commands := map[string]string{
		"java":   "-cp",
		"go":     "",
		"c":      "",
		"kotlin": "-jar",
		"python": "",
	}
	runner := Runner{
		Name:      runners[lang],
		Language:  lang,
		Command:   commands[lang],
		InputPath: inputPath,
	}

	return &runner, nil
}

// Run code runner
func (r *Runner) Run() (string, error) {
	cmd, args := buildCommand(r)
	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("running error: %v %v: %v: %s", r.Name, strings.Join(args, " "), err, output)
	}

	return string(output), nil
}

func buildCommand(r *Runner) (*exec.Cmd, []string) {
	var command *exec.Cmd
	var args []string

	switch r.Language {
	case "java":
		args = []string{r.Command, filepath.Dir(r.InputPath), "Main"}
		args = append(args, r.Extra...)
		command = exec.Command(r.Name, args...)
	case "kotlin":
		args = []string{r.Command, r.InputPath}
		args = append(args, r.Extra...)
		command = exec.Command(r.Name, args...)
	case "python":
		command = exec.Command(r.Name, r.InputPath)
	default:
		command = exec.Command(r.InputPath)
	}

	return command, args
}
