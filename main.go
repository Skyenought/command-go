package main

import (
	"flag"
	"log"
)

var name string
func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go Language", "help message")
	javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
	javaCmd.StringVar(&name, "n", "java language", "help message")
	
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "java":
		_ = javaCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}
