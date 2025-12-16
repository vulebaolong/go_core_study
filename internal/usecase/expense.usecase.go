package usecase

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-core-study/internal/model"
	"os"
	"strconv"
	"strings"
)

type ExpenseUseCase struct {
	storeFile string
}

func NewExpenseUseCase() *ExpenseUseCase {
	return &ExpenseUseCase{
		storeFile: "expenses.json",
	}
}

func (e *ExpenseUseCase) AddExpense() bool {
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

func (e *ExpenseUseCase) UpdateExpense() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chi tiêu muốn cập nhật: ")
	expenseId, _ := reader.ReadString('\n')
	expenseIdString := strings.TrimSpace(expenseId)
	expenseIdInt, err := strconv.Atoi(expenseIdString)
	if err != nil {
		fmt.Println("ID không hợp lệ")
		return false
	}

	listExpens := e.ListExpense()

	var expenseFind model.Expense
	var index int = -1

	for i, a := range listExpens {
		if a.ID == expenseIdInt {
			expenseFind = a
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Không tìm thấy chi tiêu id %d\n", expenseIdInt)
		return false
	}

	// Hiển thị thông tin cũ
	fmt.Println("\n--- Thông tin hiện tại ---")
	fmt.Println("ID:", expenseFind.ID)
	fmt.Println("Danh mục:", expenseFind.Category)
	fmt.Println("Số tiền:", expenseFind.Amount)
	fmt.Println("Ghi chú:", expenseFind.Note)

	fmt.Println("\n--- Nhập thông tin chỉnh sửa (Enter để giữ nguyên) ---")

	// ===== Category =====
	fmt.Print("Danh mục mới: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)
	if category != "" {
		expenseFind.Category = category
	}

	// ===== Amount =====
	fmt.Print("Số tiền mới: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	if amountStr != "" {
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			fmt.Println("Số tiền không hợp lệ, giữ nguyên giá trị cũ")
		} else {
			expenseFind.Amount = amount
		}
	}

	// ===== Note =====
	fmt.Print("Ghi chú mới: ")
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)
	if note != "" {
		expenseFind.Note = note
	}

	// Update lại vào list
	listExpens[index] = expenseFind

	fmt.Println("\n✅ Cập nhật thành công!")

	if err := WriteExpensesToFile(e.storeFile, listExpens); err != nil {
		fmt.Println("Lỗi lưu file:", err)
		return false
	}

	return true
}

func (e *ExpenseUseCase) DeleteExpense() bool {
	const fileName = "expenses.json"

	expenses := e.ListExpense()

	reader := bufio.NewReader(os.Stdin)

	// 2. Nhập ID cần xoá
	fmt.Print("Nhập ID chi tiêu muốn xoá: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID không hợp lệ")
		return false
	}

	// 3. Tìm & xoá
	index := -1
	for i, ex := range expenses {
		if ex.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Không tìm thấy chi tiêu ID %d\n", id)
		return false
	}

	// (tuỳ chọn) confirm
	fmt.Printf("Bạn chắc chắn muốn xoá chi tiêu [%d - %s - %d]? (y/n): ",
		expenses[index].ID,
		expenses[index].Category,
		expenses[index].Amount,
	)

	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))
	if confirm != "y" {
		fmt.Println("❌ Huỷ xoá")
		return false
	}

	// Xoá khỏi slice
	expenses = append(expenses[:index], expenses[index+1:]...)

	// 4. Ghi lại file
	if err := WriteExpensesToFile(fileName, expenses); err != nil {
		fmt.Println("Lỗi lưu file:", err)
		return false
	}

	fmt.Println("✅ Xoá thành công!")

	return true
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

func WriteExpensesToFile(filename string, expenses []model.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
