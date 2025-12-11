package controller

import (
	"go-core-study/internal/common/response"
	"go-core-study/internal/interfaces"
	"go-core-study/internal/usecase"
)

type expenseController struct {
	usecase interfaces.ExpenseUseCase
}

func NewExpenseController(usecase interfaces.ExpenseUseCase) interfaces.ExpenseController {
	return &expenseController{
		usecase: usecase,
	}
}

func (e *expenseController) AddExpense() string {
	result, err := usecase.NewExpenseUseCase().AddExpense()
	return response.HandleResponse(result, err, "Thêm chi tiêu")
}
func (e *expenseController) ListExpense() string {
	result, err := usecase.NewExpenseUseCase().ListExpense()
	return response.HandleResponse(result, err, "Liệt kê chi tiêu")
}
func (e *expenseController) ListCategories() string {
	result, err := usecase.NewExpenseUseCase().ListCategories()
	return response.HandleResponse(result, err, "Liệt kê danh mục")
}
func (e *expenseController) SummaryExpense() string {
	result, err := usecase.NewExpenseUseCase().SummaryExpense()
	return response.HandleResponse(result, err, "Tổng hợp chi tiêu")
}
