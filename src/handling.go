package main
import "fmt"

func sqrtWithErrorReturn(input float64) (float64, error) {
	if f < 0{
		return 0, errors.New("Input value is below zero.")
	}
	// ... Rest of the sqrt function
}

func main(){
	result, err := sqrtWithErrorReturn(-1)
	fmt.Println("Returned error:" + err)
}
