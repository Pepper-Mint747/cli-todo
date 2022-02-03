package main

import (
	"fmt"
	"github.com/Pepper-Mint747/goproj/interacting/todo"
	"os"
	"strings"
)

//Hard-coding the file name
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	//Use the Get method to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	// Concatenate all provided arguments with a space and
	//add to the list as an item
	default:

		//Concatenate all arguments with a space
		item := strings.Join(os.Args[1:], " ")

		//Add the task
		l.Add(item)

		//Save the new List
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
