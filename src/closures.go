package main
import "fmt"
import "reflect"

var fn = func(a int) int {
	return a
}

func main() {
	fmt.Println("Type of fn:", reflect.TypeOf(fn))
	fmt.Println("fn(5):", fn(5))
}
