package usecase

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-core-study/internal/common/response"
	"go-core-study/internal/interfaces"
	"go-core-study/internal/model"
	"os"
	"strings"
)

type expenseUseCase struct {
	storeFile string
}

func NewExpenseUseCase() interfaces.ExpenseUseCase {
	return &expenseUseCase{
		storeFile: "expenses.json",
	}
}

func (e *expenseUseCase) AddExpense() (any, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập danh mục: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Nhập số tiền: ")
	var amount int
	fmt.Scanf("%d\n", &amount)

	fmt.Print("Nhập ghi chú: ")
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)

	newExpense := model.Expense{
		Category: category,
		Amount:   amount,
		Note:     note,
	}

	var expenses []model.Expense
	data, err := os.ReadFile(e.storeFile)
	if err != nil {
		return nil, response.NewAppError(`Người dùng không tồn tại, vui lòng đăng ký`)
	}

	json.Unmarshal(data, &expenses)

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
		return nil, response.NewAppError(err.Error())
	}

	err = os.WriteFile(e.storeFile, data, 0644)
	if err != nil {
		return nil, response.NewAppError(err.Error())
	}

	return true, nil
}

func (e *expenseUseCase) ListExpense() ([]model.Expense, error) {
	// Load danh sách chi tiêu hiện tại
	var expenses []model.Expense

	data, err := os.ReadFile(e.storeFile)
	if err != nil {
		return nil, response.NewAppError(err.Error())
	}

	json.Unmarshal(data, &expenses)

	return expenses, nil
}

func (e *expenseUseCase) SummaryExpense() (any, error) {
	categories, err := e.ListCategories()
	if err != nil {
		return nil, response.NewAppError(err.Error())
	}

	if len(categories) == 0 {
		fmt.Println("Chưa có danh mục nào!")
		return []model.Expense{}, nil
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
		return []model.Expense{}, nil
	}

	category := categories[idx-1]

	expenses, err := e.ListExpense()
	if err != nil {
		return nil, response.NewAppError(err.Error())
	}
	reusltExpenses := []model.Expense{}

	for _, categoryFound := range expenses {
		if categoryFound.Category == category {
			reusltExpenses = append(reusltExpenses, categoryFound)
		}
	}

	return reusltExpenses, nil
}

func (e *expenseUseCase) ListCategories() ([]string, error) {
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
	return categories, nil
}
