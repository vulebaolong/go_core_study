# Go Core Study - Quản lý chi tiêu

## Giới thiệu
Đây là dự án mẫu giúp học viên thực hành các kiến thức Go cơ bản thông qua case study quản lý chi tiêu trên giao diện console (CLI).

## Yêu cầu hệ thống
- Go >= 1.18
- Git
- Hệ điều hành: Windows, macOS, hoặc Linux

## Hướng dẫn cài đặt và chạy dự án

### 1. Clone source code

```bash
git clone https://github.com/your-username/go-core-study.git
cd go-core-study
```

### 2. Khởi tạo Go module (nếu chưa có)

```bash
go mod init go-core-study
go mod tidy
```

### 3. Chạy ứng dụng

```bash
go run cmd/main.go
```

### 4. Cấu trúc thư mục

```
cmd/                # Chứa file main.go (entrypoint)
internal/
  controller/       # Xử lý logic điều phối
  delivery/         # Giao tiếp với user (CLI)
  di/               # Khởi tạo và kết nối các thành phần
  usecase/          # Chứa business logic
  common/response/  # Định nghĩa response chuẩn
```

## Tính năng
- Hiển thị menu quản lý chi tiêu trên console
- Thêm chi tiêu (demo)
- Liệt kê chi tiêu (demo)
- Tổng hợp chi tiêu (demo)

## Kiến thức Go thực hành
- Packages, Imports, Exported names
- Functions
- Variables, Short variable declarations, Type inference
- Basic types, Zero values
- Structs, Struct Fields, Pointer receivers, Methods
- Slices
- Switch, If, For

## Đóng góp
Mọi đóng góp, chỉnh sửa vui lòng gửi pull request hoặc liên hệ qua email.

---

**Chúc bạn học Go vui vẻ!**
# go_core_study
