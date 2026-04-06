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
-   Struct là gì trong Go?
    -   Struct là kiểu dữ liệu do bạn tự định nghĩa để gom nhiều trường (field) liên quan vào cùng một “đối tượng”.
    -   Nếu int, string là “1 giá trị đơn”
    -   thì struct là “một gói nhiều giá trị có liên quan”
    -   Ví dụ: “Chi tiêu” không chỉ có 1 số tiền, mà còn có id, danh mục, ghi chú → phù hợp dùng struct.

-   ID int
    -   Tên field là ID
    -   Kiểu dữ liệu là int
    -   Dùng để định danh mỗi khoản chi
    -   📌 Lưu ý: trong Go, field viết hoa (ID, Category, …) nghĩa là exported → code ở package khác có thể truy cập (e.ID), và các thư viện như JSON cũng “nhìn thấy” được.
    -   Nếu bạn viết: id int
        -   thì field là private (unexported) → package khác không truy cập được, JSON marshal cũng thường không encode field đó.

-   “Tag” trong struct là gì? (phần json:"...")
    -   ID int `json:"id"`
    -   Phần trong dấu backtick `...` gọi là struct tag
    -   Nó là “metadata” để thư viện khác hiểu cách xử lý field.
    -   Ở đây là tag json, nghĩa là:
    -   khi convert struct → JSON, field ID sẽ có key là "id" (chữ thường)
    -   Ví dụ:
        ```go
        e := Expense{ID: 1, Category: "Food", Amount: 50000, Note: "Lunch"}
        ```
        Khi encode JSON sẽ ra dạng:
        ```json
        {
            "id": 1,
            "category": "Food",
            "amount": 50000,
            "note": "Lunch"
        }
        ```
        📌 Nếu không có tag thì mặc định JSON sẽ dùng tên field y chang ("ID", "Category"), thường không đẹp và không theo quy ước API.
    -   omitempty: không xuất field nếu rỗng
        ```go
        Note string `json:"note,omitempty"`
        ```
        Nếu Note == "" thì JSON sẽ không có key note.
    -   "-": ẩn field
        ```go
        Password string `json:"-"`
        ```
        Password vẫn tồn tại trong struct, nhưng:
        -   Không xuất hiện trong JSON
        -   Không bị lộ khi trả API response

### 2: Tạo store quản lý danh sách chi tiêu
Tạo file internal/repository/expense_store.go:
```go
package repository

import (
    "encoding/json"
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
	return os.WriteFile(storeFile, data, 0644)
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