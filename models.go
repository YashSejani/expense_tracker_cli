package main

import(
	"time"
)

type Expense struct {
	Id int `json:"id"`
	Date time.Time `json:"date"`
	Description string `json:"description"`
	Amount int `json:"amount"`
	Category string `json:"category,omitempty"`
}

type MonthlyBudget struct{
	Month int `json:"month"`
	Year int `json:"year"`
	Budget int `json:"budget"`
}

type StorageSchema struct {
	NextId int `json:"next_id"`
	Expenses []Expense `json:"expenses"`
	Budgets []MonthlyBudget `json:"budgets"`
}

type ExpenseManager struct {
	Filename string 
	Storage StorageSchema
}