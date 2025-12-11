Multiple results (trả về nhiều giá trị cùng lúc)  
Named return values  
Constants  
Arrays, Maps, Range  

## step 1
### 1: Định nghĩa struct cho chi tiêu
Tạo file internal/model/expense.go:
```go
package model

type Expense struct {
    ID       int    `json:"id"`
    Category string `json:"category"`
    Amount   int    `json:"amount"`
    Note     string `json:"note"`
}
```

### 2: Tạo store quản lý danh sách chi tiêu
Tạo file internal/repository/expense_store.go:
```go
package repository

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "go-core-study/internal/model"
)

// 3. Constants:
const storeFile = "expenses.json"

func SaveExpenses(expenses []model.Expense) error {
    data, err := json.MarshalIndent(expenses, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(storeFile, data, 0644)
}

// 1. Multiple results (trả về nhiều giá trị cùng lúc):
// 2. Named return values:
func LoadExpenses() ([]model.Expense, error) {
    var expenses []model.Expense
    file, err := os.Open(storeFile)
    if err != nil {
        if os.IsNotExist(err) {
            return expenses, nil // Trả về slice rỗng nếu chưa có file
        }
        return nil, err
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(data, &expenses)
    return expenses, err
}
```

### 3: Sử dụng store trong usecase/controller
Khi thêm chi tiêu mới: đọc danh sách, thêm phần tử, ghi lại vào file.
Khi liệt kê chi tiêu: đọc từ file và trả về.

### 4: Kiểm tra hoạt động
Thêm chi tiêu mới sẽ lưu vào file expenses.json.
Liệt kê chi tiêu sẽ đọc từ file và hiển thị ra CLI.

### Tóm tắt các bước:

Định nghĩa struct Expense.
Tạo hàm lưu/đọc danh sách chi tiêu từ file JSON.
Sử dụng các hàm này trong usecase/controller.
Kiểm tra file expenses.json được tạo và cập nhật đúng.