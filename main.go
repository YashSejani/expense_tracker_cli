package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No command provided.")
		PrintGeneralHelp()
		return
	}

	manager := ExpenseManager{
		Filename: "expense.json",
	}

	manager.Load()

	switch os.Args[1] {

	case "help":
		PrintGeneralHelp()

	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Description of the expense")
		amount := addCmd.Int("amount", 0, "Amount of the expense")
		addCmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Error: Both --description and a valid positive --amount are required.")
			fmt.Println("Example: expense-tracker add --description \"Lunch\" --amount 20")
			return
		}
		manager.AddExpense(*description, *amount)

	case "list":
		manager.ListExpense()

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := delCmd.Int("id", 0, "Id of expense")
		delCmd.Parse(os.Args[2:])

		if *id <= 0 {
			fmt.Println("Error: A valid positive --id is required.")
			fmt.Println("Example: expense-tracker delete --id 1")
			return
		}
		manager.DeleteExpense(*id)

	case "summary":
		summCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summCmd.Int("month", -1, "month of expense (1-12)")
		year := summCmd.Int("year", -1, "year of expense")
		summCmd.Parse(os.Args[2:])

		if *month != -1 && (*month < 1 || *month > 12) {
			fmt.Println("Error: --month must be a value between 1 and 12.")
			return
		}
		manager.GetSummary(*month, *year)

	default:
		fmt.Printf("Error: Unknown command '%s'.\n", os.Args[1])
		PrintGeneralHelp()
		return
	}

	// Only save if the commands executed didn't return early with failures
	manager.Save()
}