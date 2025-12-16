package controller

import (
	"go-core-study/internal/common/response"
	"go-core-study/internal/usecase"
)

type ExpenseController struct {
	usecase *usecase.ExpenseUseCase
}

func NewExpenseController(usecase *usecase.ExpenseUseCase) *ExpenseController {
	return &ExpenseController{
		usecase: usecase,
	}
}

func (e *ExpenseController) AddExpense() string {
	result := usecase.NewExpenseUseCase().AddExpense()
	return response.NewAppSuccess(result, "Thêm chi tiêu")
}
func (e *ExpenseController) ListExpense() string {
	result := usecase.NewExpenseUseCase().ListExpense()
	return response.NewAppSuccess(result, "Liệt kê chi tiêu")
}
func (e *ExpenseController) UpdateExpense() string {
	result := usecase.NewExpenseUseCase().UpdateExpense()
	return response.NewAppSuccess(result, "Cập nhật kê chi tiêu")
}
func (e *ExpenseController) DeleteExpense() string {
	result := usecase.NewExpenseUseCase().DeleteExpense()
	return response.NewAppSuccess(result, "Xoá chi tiêu")
}
func (e *ExpenseController) SummaryExpense() string {
	result := usecase.NewExpenseUseCase().SummaryExpense()
	return response.NewAppSuccess(result, "Tổng hợp chi tiêu")
}
