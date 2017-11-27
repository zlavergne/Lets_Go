package main
import "fmt"

func divide(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func main() {
	q, r := divide(11, 3)
	fmt.Println("divide(11/3):")
	fmt.Println("quotient", q)
	fmt.Println("remainder", r)
}
