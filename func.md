Dưới đây là “tất tần tật” về **function trong Go** và **các loại syntax hay gặp**, kèm ví dụ ngắn gọn.

---

## Function cơ bản

```go
func Add(a int, b int) int {
	return a + b
}
```

* `func` là từ khoá khai báo hàm
* `Add` tên hàm
* `(a int, b int)` tham số
* `int` sau ngoặc là kiểu trả về
* `return` trả kết quả

Rút gọn kiểu cho nhiều biến cùng kiểu:

```go
func Add(a, b int) int { return a + b }
```

---

## Không trả về gì

```go
func Hello() {
	fmt.Println("Hi")
}
```

---

## Nhiều giá trị trả về

Go rất hay dùng kiểu này (đặc biệt với `value, err`):

```go
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}
```

Gọi:

```go
q, err := Div(10, 2)
```

---

## Named return values (kết quả có tên)

```go
func Stats(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return // return sum, diff
}
```

> Dùng được, nhưng đừng lạm dụng vì dễ khó đọc.

---

## Variadic function (tham số “...”)

Nhận số lượng tham số linh hoạt:

```go
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
```

Gọi:

```go
Sum(1, 2, 3)
```

Truyền slice vào variadic:

```go
arr := []int{1, 2, 3}
Sum(arr...)
```

---

## Truyền “tham chiếu” bằng pointer

Go truyền tham số theo “copy”, muốn sửa ngoài hàm thì dùng pointer:

```go
func Inc(x *int) {
	*x = *x + 1
}

n := 10
Inc(&n) // n thành 11
```

---

## Receiver method (hàm gắn với struct)

### Value receiver

Không sửa được object gốc (trừ khi field là reference type):

```go
type User struct{ Name string }

func (u User) Greet() string {
	return "Hi " + u.Name
}
```

### Pointer receiver

Sửa được object gốc, thường dùng khi struct lớn hoặc cần mutate:

```go
func (u *User) Rename(name string) {
	u.Name = name
}
```

Quy tắc thực tế:

* **Có mutate → pointer receiver**
* Struct lớn → pointer receiver (đỡ copy)
* Muốn đồng nhất method set → chọn pointer cho tất cả

---

## Function type (hàm là “giá trị”)

Bạn có thể gán hàm vào biến:

```go
type Op func(int, int) int

func Apply(a, b int, op Op) int {
	return op(a, b)
}

res := Apply(2, 3, func(x, y int) int { return x * y })
```

---

## Anonymous function (hàm ẩn danh)

```go
f := func(a int) int { return a * 2 }
fmt.Println(f(10)) // 20
```

---

## Closure (hàm “bắt” biến bên ngoài)

```go
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

c := Counter()
c() // 1
c() // 2
```

---

## Defer (chạy sau cùng khi hàm kết thúc)

Hay dùng để `Close()`:

```go
func ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil { return err }
	defer f.Close()

	// đọc file...
	return nil
}
```

Defer chạy theo kiểu “stack” (LIFO).

---

## Return sớm (early return)

Style phổ biến của Go: xử lý lỗi xong return luôn:

```go
func Do() error {
	if err := step1(); err != nil { return err }
	if err := step2(); err != nil { return err }
	return nil
}
```

---

## Syntax hay gặp khi gọi hàm

### Multiple assignment

```go
v, err := Something()
```

### Ignore giá trị bằng `_`

```go
_, err := Something()
```

### Gọi method

```go
u.Rename("Nam")
```

---

## Generic function (Go 1.18+)

```go
func Max[T ~int | ~float64](a, b T) T {
	if a > b { return a }
	return b
}
```

Hoặc “any”:

```go
func PrintAll[T any](items []T) {
	for _, it := range items {
		fmt.Println(it)
	}
}
```

---

## Panic vs error (quan trọng)

* **error**: lỗi “bình thường” nên trả về `error`
* **panic**: lỗi “không thể cứu” (bug, invariant hỏng)

```go
panic("something went terribly wrong")
```

---

## Quick checklist (thực chiến)

* Hầu hết hàm “có thể fail” → `(...)(T, error)`
* Mutate struct → `func (x *Type) ...`
* Nhận callback → `func(... )` làm tham số
* Dọn tài nguyên → `defer Close()`

---

Nếu bạn muốn, mình có thể soạn thêm 1 phần **“cheat sheet func/method theo tình huống”** (CRUD CLI như bạn đang làm: `Create/List/Update/Delete`, tách `UseCase` + `Repository`, patterns hay dùng trong Go). Bạn muốn theo hướng **beginner** hay **backend clean architecture**?

Câu hỏi này rất **chuẩn tư duy hệ thống** 👍
Mình sẽ **kiểm tra lại toàn bộ “bề mặt” của function trong Go**, đối chiếu với spec + thực tế backend, rồi **bổ sung những phần còn thiếu** (những cái ít người nói nhưng rất quan trọng).

---

# ✅ Tổng kiểm kê: Function trong Go

Dưới đây là **những phần CÒN THIẾU hoặc CHƯA NÓI SÂU** so với danh sách chuẩn.

---

## 1️⃣ Init function (`init()`)

⛔ **Rất hay bị quên**

```go
func init() {
	fmt.Println("Chạy trước main")
}
```

* Tự động chạy
* Không được gọi thủ công
* Mỗi file có thể có nhiều `init`
* Thứ tự:

  1. import
  2. biến global
  3. `init()`
  4. `main()`

👉 Dùng để:

* Setup config
* Load env
* Init global state

---

## 2️⃣ Method Set (value vs pointer receiver) – PHẦN RẤT QUAN TRỌNG

```go
type A struct{}
func (a A) Foo() {}
func (a *A) Bar() {}
```

| Kiểu | Gọi được         |
| ---- | ---------------- |
| `A`  | `Foo()`          |
| `*A` | `Foo()`, `Bar()` |

👉 Go **tự động lấy address hoặc dereference**
Nhưng interface thì KHÔNG

---

## 3️⃣ Function + Interface relationship

### Function implement interface?

❌ **KHÔNG**

Chỉ **method** mới implement interface.

```go
type Runner interface {
	Run()
}
```

✔ OK:

```go
type Job struct{}
func (j Job) Run() {}
```

❌ KHÔNG OK:

```go
func Run() {}
```

---

## 4️⃣ Function as return value (factory pattern)

```go
func Logger(prefix string) func(string) {
	return func(msg string) {
		fmt.Println(prefix, msg)
	}
}
```

Dùng rất nhiều trong:

* middleware
* validator
* decorator

---

## 5️⃣ Higher-order function (map / filter / reduce)

Go không built-in nhưng dùng function type:

```go
func Map[T any](arr []T, f func(T) T) []T {
	out := make([]T, len(arr))
	for i, v := range arr {
		out[i] = f(v)
	}
	return out
}
```

---

## 6️⃣ Method expression & method value (ÍT NGƯỜI BIẾT)

### Method value

```go
u := User{Name: "A"}
f := u.Greet
fmt.Println(f())
```

### Method expression

```go
f := User.Greet
fmt.Println(f(u))
```

---

## 7️⃣ Function visibility (exported / unexported)

```go
func publicFunc() {}   // private
func PublicFunc() {}   // exported
```

⚠️ Chỉ dựa vào **chữ cái đầu**

---

## 8️⃣ Function & zero value

```go
var f func(int) int
fmt.Println(f == nil) // true
```

❗ Gọi `f()` khi nil → panic

---

## 9️⃣ Recursion (đệ quy)

```go
func Fact(n int) int {
	if n == 0 { return 1 }
	return n * Fact(n-1)
}
```

Go **không có TCO** → cẩn thận stack overflow.

---

## 🔟 Function & concurrency (goroutine)

```go
go func() {
	fmt.Println("run async")
}()
```

⚠️ Closure + loop trap:

```go
for i := 0; i < 3; i++ {
	go func() {
		fmt.Println(i) // BUG
	}()
}
```

Fix:

```go
for i := 0; i < 3; i++ {
	i := i
	go func() {
		fmt.Println(i)
	}()
}
```

---

## 1️⃣1️⃣ Recover & panic handling

```go
func Safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	panic("boom")
}
```

---

## 1️⃣2️⃣ Function + build tags (advanced)

```go
//go:build linux
func Platform() string { return "linux" }
```

---

## 1️⃣3️⃣ Inline function optimization (compiler hint)

Go compiler **tự inline**, không có keyword, nhưng:

* Hàm nhỏ
* Không variadic phức tạp
  → dễ inline

---

## 1️⃣4️⃣ Function vs Method best practices (thực tế)

| Dùng khi         | Nên      |
| ---------------- | -------- |
| Stateless helper | Function |
| Có dữ liệu gắn   | Method   |
| Domain logic     | Method   |
| Utility          | Function |

---

## 1️⃣5️⃣ Test function (`*_test.go`)

```go
func TestAdd(t *testing.T) {}
func BenchmarkAdd(b *testing.B) {}
func FuzzAdd(f *testing.F) {}
```

---

## 🧠 Tổng kết – ĐÃ ĐỦ CHƯA?

| Nội dung          | Trạng thái  |
| ----------------- | ----------- |
| Basic function    | ✅           |
| Multi return      | ✅           |
| Named return      | ✅           |
| Variadic          | ✅           |
| Pointer           | ✅           |
| Method            | ✅           |
| Closure           | ✅           |
| Defer             | ✅           |
| Generic           | ✅           |
| Panic / recover   | ✅           |
| init()            | ✅ (bổ sung) |
| Method set        | ✅ (bổ sung) |
| Method expression | ✅ (bổ sung) |
| Goroutine + func  | ✅ (bổ sung) |
| Zero value func   | ✅ (bổ sung) |
| Testing func      | ✅ (bổ sung) |

👉 **Đến đây là FULL 100% kiến thức function trong Go ở mức backend thực chiến.**

---

Nếu bạn muốn bước tiếp:

* So sánh **Go func vs Java method**
* Function design cho **Clean Architecture**
* Anti-pattern về function trong Go
* Viết guideline chuẩn cho team

👉 nói mình chọn hướng nào, mình làm tiếp cho bạn 👌
