package main

import (
	"errors"
	"expense-tracker/expense"
	"fmt"
	"os"
)

func handleErrors(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const filename = "expense.txt"

func main() {
	if len(os.Args) < 2 {
		handleErrors(errors.New("commands should contains more that 2 argument"))
	}
	action := os.Args[1]

	switch action {
	case "add":
		if err := expense.Add(filename); err != nil {
			handleErrors(err)
		}
	case "update":
		if err := expense.Update(filename); err != nil {
			handleErrors(err)
		}
	case "delete":
		if err := expense.Delete(filename); err != nil {
			handleErrors(err)
		}
	case "list":
		if err := expense.List(filename); err != nil {
			handleErrors(err)
		}
	case "summery":
		expense.Summery(filename)
	default:
		fmt.Println("You dont have any expenses")
	}
}
