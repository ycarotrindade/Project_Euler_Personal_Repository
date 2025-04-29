package main
import "fmt"

func main() {
	var max int
	var sum int
	fmt.Println("Type the max value:")
	fmt.Scan(&max)
	for i := 1; i<max; i++{
		if verifyMult(i){
			sum += i
		}
	}
	fmt.Printf("The sum of the multiples of 3 or 5 below \033[32m %d \033[0m is \033[32m %d \033[0m",max,sum)
}

func verifyMult(number int) bool{
	if number % 3 == 0 || number % 5 == 0{
		return true
	}else{
		return false
	}
}