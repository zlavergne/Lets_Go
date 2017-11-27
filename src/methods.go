package main
import "fmt"

type cube struct {
	length int
	width int
	height int
}

var a = cube{5,5,5}

func (c cube) volume() int {
	return c.length * c.width * c.height
}

func main() {
	fmt.Println("a.volume():", a.volume())
}
