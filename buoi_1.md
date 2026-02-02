
Packages, Imports, Exported names
Functions
Variables, Short variable declarations, Type inference
Basic types, Zero values
Structs, Struct Fields, Pointer receivers, Methods
Slices
Switch, If, For

## step 1
Khởi tạo và giải thích cấu trúc cơ bản

### 1. Khởi tạo project Go
Mở terminal tại thư mục bạn muốn tạo project, chạy:  
```bash
go mod init go-core-study
```
-   Lệnh sẽ tạo ra file go.mod để quản lý module và dependencies.
-   Gồm:
    -   Tên project
    -   Phiên bản Go
    -   Thư viện phụ thuộc (dependencies)
-   File go.mod giống như:
    -   package.json (Node.js)
    -   pom.xml (Java Maven)
    -   requirements.txt (Python)

### 2.  Chạy Hello world
-   Tạo file main.go
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello world")
    }
    ```
    -   main.go không bắt buộc phải tên là main.go, bạn có thể đặt app.go, server.go…
    -   Quan trọng là file đó thuộc package main và có hàm func main().
    -   main là tên hàm đặc biệt: Go runtime sẽ gọi hàm này đầu tiên khi chạy chương trình.
    -   main() không nhận tham số và không trả về gì.
    -   Go build theo “package”, không theo “file”
        -   Trong 1 thư mục (ví dụ cmd/app/), bạn có thể có nhiều file cùng package:
            -   main.go
            -   router.go
            -   handler.go
        -   Khi xử dụng hàm giữa các file thì không cần phải import vì chúng cùng package
        
### 2. Tạo cấu trúc thư mục
Tạo các folder như sau:  
```bash
go-core-study/
│
├── cmd/                # Chứa file main.go (entrypoint)
│   └── main.go
│
├── internal/           # Chứa toàn bộ code logic, chia nhỏ theo domain
│   ├── controller/     # Xử lý logic điều phối, nhận request từ delivery, gọi usecase
│   ├── delivery/       # Giao tiếp với user (CLI, HTTP, v.v.)
│   ├── di/             # Dependency Injection, khởi tạo và kết nối các thành phần
│   ├── model/          # Định nghĩa các struct dữ liệu (nếu có)
│   ├── repository/     # Tầng truy xuất dữ liệu (nếu có)
│   ├── usecase/        # Chứa business logic (nghiệp vụ)
│   └── common/         # Chứa các thành phần dùng chung (ví dụ: response)
│       └── response/
│
├── go.mod
└── road_map.md         # (Tùy chọn) Ghi chú roadmap phát triển
```

### 3. Giải thích nhiệm vụ từng folder
cmd/: Chứa file main.go, là entrypoint chạy chương trình.
internal/: Chứa toàn bộ code logic, chia nhỏ theo domain, không export ra ngoài module.
controller/: Xử lý logic điều phối, nhận request từ delivery, gọi usecase.
delivery/: Giao tiếp với user, ví dụ CLI, HTTP, gRPC, v.v.
di/: Dependency Injection, khởi tạo và kết nối các thành phần lại với nhau.
model/: Định nghĩa các struct dữ liệu (nếu có).
usecase/: Chứa business logic (nghiệp vụ chính).
common/: Chứa các thành phần dùng chung, ví dụ response, error, helper.
response/: Định nghĩa struct và hàm trả về response chuẩn.

## step 2
Hoành thành giao diện CLI

### Bước 1: Tạo file CLI
Tạo file cli.delivery.go với nội dung khởi tạo giao diện CLI:
```go
package delivery

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

type CLI struct{}

func NewCLI() *CLI {
    return &CLI{}
}

func (cli *CLI) Run() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n====== Quản lý chi tiêu ======")
        fmt.Println("1) Thêm chi tiêu")
        fmt.Println("2) Liệt kê chi tiêu")
        fmt.Println("3) Tổng hợp theo danh mục")
        fmt.Println("0) Thoát")
        fmt.Print(">>Chọn: ")

        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            log.Println("Thêm chi tiêu (chưa xử lý)")
        case "2":
            log.Println("Liệt kê chi tiêu (chưa xử lý)")
        case "3":
            log.Println("Tổng hợp theo danh mục (chưa xử lý)")
        case "0":
            fmt.Println("Tạm biệt!")
            return
        default:
            fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
        }
    }
}
```

### Bước 2: Sửa file main.go để gọi CLI
Sửa file main.go để khởi tạo và chạy CLI:

### Bước 3: Chạy thử ứng dụng
Mở terminal tại thư mục gốc dự án, chạy:

Bạn sẽ thấy giao diện CLI hiện ra, chọn các menu sẽ log ra thông báo tương ứng.

### Tóm tắt:
Tạo struct CLI và hàm Run để hiện menu.
Chỉ log ra thông báo, chưa xử lý logic.
Sửa main.go để gọi CLI.
Chạy thử để kiểm tra giao diện console.

