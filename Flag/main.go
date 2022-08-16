package main

import (
	"flag"
	"fmt"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myFlags arrayFlags

func main() {
	flag.Var(&myFlags, "list1", "Some description for this param.")
	flag.Parse()
	for _, i := range myFlags {
		fmt.Println(i)
	}
}

func showArgs() {
	for i, args := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, args)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
