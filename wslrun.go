/*
Helper for run linux utilities in windows. It converts windows paths to wsl paths

For example
	wslrun diff --color c:\projects\hello\index.js c:\projects\oldHello\index.js
	will launched in wsl as
	diff --color /mnt/c/projects/hello/index.js /mnt/c/projects/oldHello/index.js
*/
package main

import (
	"github.com/mattn/go-colorable"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var isPath = regexp.MustCompile(`^[a-zA-Z]:\\`)

func main() {
	var args []string
	for _, arg := range os.Args[1:] {
		if isPath.MatchString(arg) {
			arg = "/mnt/" + arg[0:1] + "/" + strings.Replace(arg[3:], `\`, `/`, -1)
		}
		args = append(args, arg)
	}

	cmd := exec.Command("wsl", "script", "-qfc", strings.Join(args, " "), "/dev/null")
	cmd.Stdout = colorable.NewColorableStdout()

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
