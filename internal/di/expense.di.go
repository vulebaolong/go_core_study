package di

import (
	"go-core-study/internal/controller"
	"go-core-study/internal/delivery"
	"go-core-study/internal/usecase"
)

func NewAppContainer() *delivery.CLI {
	usecase := usecase.NewExpenseUseCase()
	controller := controller.NewExpenseController(usecase)
	cli := delivery.NewCLI(controller)

	return cli
}
