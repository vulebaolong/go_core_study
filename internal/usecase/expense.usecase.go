package usecase

type ExpenseUseCase struct{}

func NewExpenseUseCase() *ExpenseUseCase {
	return &ExpenseUseCase{}
}

func (e *ExpenseUseCase) AddExpense() bool {
	return true
}

func (e *ExpenseUseCase) ListExpense() []string {
	return []string{
		"Food - 20000",
		"Coffee - 15000",
	}
}

func (e *ExpenseUseCase) UpdateExpense() []string {
	return []string{
		"Food - 20000",
		"Coffee - 15000",
	}
}

func (e *ExpenseUseCase) DeleteExpense() []string {
	return []string{
		"Food - 20000",
		"Coffee - 15000",
	}
}

func (e *ExpenseUseCase) SummaryExpense() []string {
	return []string{
		"Coffee - 15000",
		"Food - 20000",
	}
}
