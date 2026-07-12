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

