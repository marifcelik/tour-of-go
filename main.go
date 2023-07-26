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
	var list []string
	var result string
	var watch *exec.Cmd
	var n bool

	flag.BoolVar(&n, "n", false, "do not use watch mode")
	flag.Parse()

	dirs, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dirs {
		if v.IsDir() && !strings.HasPrefix(v.Name(), ".") {
			list = append(list, v.Name())
		}
	}

	prompt := &survey.Select{Message: "choose dir", Options: list}
	errSurvey := survey.AskOne(prompt, &result)
	if errSurvey != nil {
		if errSurvey.Error() == "interrupt" {
			log.Fatal("interrupted")
		}
	}

	if n {
		watch = exec.Command("go", "run", ".")
		fmt.Printf("running %s folder\n", result)
	} else {
		watch = exec.Command("watchexec", "-c", "-e go", "go run .")
	}

	watch.Dir = result
	watch.Stdout = os.Stdout
	watch.Stderr = os.Stderr
	watch.Run()

}
