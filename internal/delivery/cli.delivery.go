package delivery

import (
	"bufio"
	"fmt"
	"go-core-study/internal/controller"
	"log"
	"os"
	"strings"
)

type CLI struct {
	controller *controller.ExpenseController
}

func NewCLI(controller *controller.ExpenseController) *CLI {
	return &CLI{
		controller: controller,
	}
}

func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n====== Quản lý chi tiêu ======")
		fmt.Println("1) Thêm chi tiêu")
		fmt.Println("2) Liệt kê chi tiêu")
		fmt.Println("3) Sửa chi tiêu")
		fmt.Println("4) Xoá chi tiêu")
		fmt.Println("5) Tổng hợp theo danh mục")
		fmt.Println("0) Thoát")
		fmt.Println("")
		fmt.Print(">>Chọn: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			result := cli.controller.AddExpense()
			log.Printf("%+v\n", result)
		case "2":
			result := cli.controller.ListExpense()
			log.Printf("%+v\n", result)
		case "3":
			result := cli.controller.UpdateExpense()
			log.Printf("%+v\n", result)
		case "4":
			result := cli.controller.DeleteExpense()
			log.Printf("%+v\n", result)
		case "5":
			result := cli.controller.SummaryExpense()
			log.Printf("%+v\n", result)
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
	}
}
