package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No command provided. Run 'expense-tracker help' to see the layout structure.")
		return
	}

	manager := ExpenseManager {
		Filename: "expense.json",
	}

	manager.Load()

	switch(os.Args[1]) {

	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)

		description := addCmd.String("description", "", "Description of the expense")
		amount := addCmd.Int("amount", 0, "Amount of the expense")

		addCmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Error: Both --description and a valid --amount are required.")
			return
		}

		manager.AddExpense(*description, *amount)
	
	case "delete":
		
	}
}