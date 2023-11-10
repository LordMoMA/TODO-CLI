package main

import (
	"flag"
	"fmt"
	todo "github.com/LordMoMA/TODO-CLI"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	task := flag.String("task", "", "Task to be included in the TODO list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Smart TODO App developed for work efficiency:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2023\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information:")
		flag.PrintDefaults()
	}
	flag.Parse()

	l := &todo.List{}

	if err := l.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, t := range *l {
			if !t.Done {
				fmt.Println(t.Task)
			}
		}
	case complete != nil && *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case task != nil && *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid option")
		os.Exit(1)
	}
}
