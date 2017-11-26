package main
import "fmt"
import "reflect"

var a = struct{
	length int
	width int
	height int
}{5,5, 5}

func main() {
	fmt.Println("Type of a:", reflect.TypeOf(a))
	fmt.Println("a:", a)
}
