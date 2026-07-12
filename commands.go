package main

import (
	"time"
	"fmt"
	"slices"
)

func (em *ExpenseManager) AddExpense(UserDescription string, Amount int) {
	newExpense := Expense{
		Id:          em.Storage.NextId,
		Date:        time.Now(),
		Description: UserDescription,
		Amount:      Amount,
	}
	em.Storage.Expenses = append(em.Storage.Expenses, newExpense)
	fmt.Printf("Expense added successfully (ID: %d)",em.Storage.NextId)
	em.Storage.NextId++
}

func (em *ExpenseManager) ListExpense() {
	if len(em.Storage.Expenses) == 0 {
		fmt.Println("No expenses found. Use 'add' to log your first expense!")
		return
	}

	fmt.Println("ID\tDate\t\tDescription\tAmount")
	fmt.Println("-----------------------------------------------")
	for _, expense := range em.Storage.Expenses {
		dateStr := expense.Date.Format("2006-01-02")
		fmt.Printf("%d\t%s\t%s\t\t$%d\n", expense.Id, dateStr, expense.Description, expense.Amount)
	}
}

func (em *ExpenseManager) DeleteExpense(id int) {
	for i, expense := range em.Storage.Expenses {
		if(expense.Id == id) {
			em.Storage.Expenses = slices.Delete(em.Storage.Expenses, i, i+1)
			fmt.Printf("Expense deleted successfully (ID: %d)\n", id)
			return
		}
	}
	fmt.Printf("Error: Expense with ID %d does not exist.\n", id)
}

func (em *ExpenseManager) GetSummary(month, year int) {
	Total := 0
	for _, expense := range em.Storage.Expenses {
		matchMonth := (month == -1 || int(expense.Date.Month()) == month)
		matchYear := (year == -1 || expense.Date.Year() == year)

		if matchMonth && matchYear {
			Total += expense.Amount
		}
	}

	if month != -1 {
		monthName := time.Month(month).String()
		fmt.Printf("Total expenses for %s: $%d\n", monthName, Total)
	} else {
		fmt.Printf("Total expenses: $%d\n", Total)
	}
}

func PrintGeneralHelp() {
	fmt.Println("====================================================================")
	fmt.Println("\t\tEXPENSE TRACKER CLI APPLICATION")
	fmt.Println("====================================================================")
	fmt.Println("\nUsage:")
	fmt.Println("  expense-tracker <command> [flags]")

	fmt.Println("\nAvailable Commands & Syntax Details:")
	fmt.Println("-----------------------------------------------------------------------")
	
	fmt.Println("  add      Log a new expense to the system.")
	fmt.Println("           Flags Required:")
	fmt.Println("             --description \"text\"  The item name or details (e.g., \"Lunch\")")
	fmt.Println("             --amount <integer>     The positive numeric cost value")
	fmt.Println("           Example: expense-tracker add --description \"Groceries\" --amount 45")
	fmt.Println()

	fmt.Println("  list     Display a tabular chronological index of all recorded expenses.")
	fmt.Println("           Flags Required: None")
	fmt.Println("           Example: expense-tracker list")
	fmt.Println()

	fmt.Println("  delete   Permanently erase a logged expense by its system ID reference.")
	fmt.Println("           Flags Required:")
	fmt.Println("             --id <integer>         The unique identification key of the entry")
	fmt.Println("           Example: expense-tracker delete --id 3")
	fmt.Println()

	fmt.Println("  summary  Compute and display accumulated financial statistics.")
	fmt.Println("           Flags Optional:")
	fmt.Println("             --month <1-12>        Filter calculations to a specific month")
	fmt.Println("             --year <integer>      Filter calculations to a specific year")
	fmt.Println("           Examples:")
	fmt.Println("             expense-tracker summary")
	fmt.Println("             expense-tracker summary --month 8")
	fmt.Println("             expense-tracker summary --month 12 --year 2026")
	fmt.Println()

	fmt.Println("  help     Render this comprehensive interactive syntax guidelines layout.")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("Pro-Tip: You can append '--help' to any specific subcommand to view")
	fmt.Println("            Go's automated flag parsing constraints.")
	fmt.Println("            Example: expense-tracker add --help")
	fmt.Println("=======================================================================")
}