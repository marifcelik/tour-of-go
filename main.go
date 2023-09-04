package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "do not use watch mode")
	flag.Parse()

	dirs, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	var list []string
	for _, v := range dirs {
		if v.IsDir() && !strings.HasPrefix(v.Name(), ".") {
			list = append(list, v.Name())
		}
	}

	var result string
	prompt := &survey.Select{Message: "choose dir", Options: list}
	errSurvey := survey.AskOne(prompt, &result)
	if errSurvey != nil {
		if errSurvey.Error() == "interrupt" {
			log.Fatal("interrupted")
		}
	}

	_, err = exec.LookPath("watchexec")
	var watch *exec.Cmd
	if n {
		watch = exec.Command("go", "run", ".")
		fmt.Printf("running %s folder\n", result)
	} else if err != nil {
		fmt.Println("watchexec not found, using go run")
		watch = exec.Command("go", "run", ".")
	} else {
		watch = exec.Command("watchexec", "-c", "-r", "-e go", "--", "go", "run", ".")
	}

	watch.Dir = result
	watch.Stdout = os.Stdout
	watch.Stderr = os.Stderr
	watch.Run()
}
