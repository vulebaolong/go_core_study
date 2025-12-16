package usecase

import (
	"encoding/json"
	"fmt"
	"go-core-study/internal/model"
	"os"
)

type ExpenseUseCase struct {
	storeFile string
}

func NewExpenseUseCase() *ExpenseUseCase {
	return &ExpenseUseCase{
		storeFile: "expenses.json",
	}
}

func (e *ExpenseUseCase) AddExpense(newExpense model.Expense) bool {
	// Load danh sách chi tiêu hiện tại
	var expenses []model.Expense
	data, err := os.ReadFile(e.storeFile)
	if err == nil {
		json.Unmarshal(data, &expenses)
	}

	// Tìm ID lớn nhất
	maxID := 0
	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}

	newExpense.ID = maxID + 1

	expenses = append(expenses, newExpense)

	// Lưu lại danh sách
	data, err = json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return false
	}

	err = os.WriteFile(e.storeFile, data, 0644)

	return err == nil
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

func (e *ExpenseUseCase) ListExpense() []model.Expense {
	// Load danh sách chi tiêu hiện tại
	var expenses []model.Expense
	data, err := os.ReadFile(e.storeFile)
	if err == nil {
		json.Unmarshal(data, &expenses)
	}

	return expenses
}

func (e *ExpenseUseCase) SummaryExpense() []model.Expense {
	// Lấy tất cả danh mục đang có
	categories := e.ListCategories()
	if len(categories) == 0 {
		fmt.Println("Chưa có danh mục nào!")
		return []model.Expense{}
	}
	fmt.Println("Các danh mục hiện có:")
	for i, cat := range categories {
		fmt.Printf("%d) %s\n", i+1, cat)
	}
	fmt.Print("Chọn số danh mục để tổng hợp: ")
	var idx int
	fmt.Scanf("%d\n", &idx)

	if idx < 1 || idx > len(categories) {
		fmt.Println("Lựa chọn không hợp lệ!")
		return []model.Expense{}
	}

	category := categories[idx-1]

	expenses := e.ListExpense()
	reusltExpenses := []model.Expense{}

	for _, categoryFound := range expenses {
		if categoryFound.Category == category {
			reusltExpenses = append(reusltExpenses, categoryFound)
		}
	}

	return reusltExpenses
}

// Liệt kê các danh mục chi tiêu đang có
func (e *ExpenseUseCase) ListCategories() []string {
	const storeFile = "expenses.json"
	var expenses []model.Expense
	var categoriesMap = make(map[string]struct{})
	var categories []string
	data, err := os.ReadFile(storeFile)
	if err == nil {
		json.Unmarshal(data, &expenses)
		for _, exp := range expenses {
			if _, exists := categoriesMap[exp.Category]; !exists {
				categoriesMap[exp.Category] = struct{}{}
				categories = append(categories, exp.Category)
			}
		}
	}
	return categories
}
