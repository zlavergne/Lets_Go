package main
import(
	"errors"
	"fmt"
)

func sqrtWithErrorReturn(input float64) (float64, error) {
	if input < 0{
		return 0, errors.New("Input value is below zero.")
	}
	// ... Rest of the sqrt function

	return 0, nil
}

func main(){
	result, err := sqrtWithErrorReturn(-1)
	fmt.Print("Returned Value: ")
	fmt.Print(result);
	fmt.Print("\nReturned Error: ")
	fmt.Print(err)
	fmt.Print("\n")
}
