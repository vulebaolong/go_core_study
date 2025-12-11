package interfaces

import "go-core-study/internal/model"

type ExpenseUseCase interface {
	AddExpense() (any, error)
	ListExpense() ([]model.Expense, error)
	SummaryExpense() (any, error)
	ListCategories() ([]string, error)
}

type ExpenseController interface {
	AddExpense() string
	ListExpense() string
	ListCategories() string
	SummaryExpense() string
}
