package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const format = "2006-01-02"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "A simple tool for creating notes ordered by date\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	var (
		dir    = flag.String("dir", ".", "Where to place the note")
		editor = flag.String("editor", "vim", "Editor to open note in")
		ext    = flag.String("ext", "md", "File extension to use")
	)
	flag.Parse()

	filePostfix := flag.Arg(0)
	if filePostfix == "" {
		log.Fatal("missing filename")
	}

	filePrefix := time.Now().Format(format)

	file := fmt.Sprintf(`%s/%s-%s.%s`, *dir, filePrefix, filePostfix, *ext)

	cmd := exec.Command(*editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
