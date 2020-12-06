package main

import (
	"code-runner/runner"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	argc := len(os.Args)

	if argc != 3 {
		fmt.Println("Error, wrong arguments.\nUsage: ./code-runner \"language\" \"input_filepath\" ")
		os.Exit(1)
	}

	r, err := runner.NewRunner(args[1], args[2])

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	out, err := r.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(out)
}
