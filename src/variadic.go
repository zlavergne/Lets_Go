package main
import "fmt"

func f(b ...int) (length int) {
	length = len(b)
	return
}

func main() {
	fmt.Println("f(0,4,9,8,23,11):", f(0,4,9,8,23,11))
}
