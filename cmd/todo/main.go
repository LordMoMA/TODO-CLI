package main

import (
	"fmt"
	todo "github.com/LordMoMA/TODO-CLI"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, t := range *l {
			fmt.Println(t.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
