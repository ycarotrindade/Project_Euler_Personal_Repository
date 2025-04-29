package main

import "fmt"

func main() {
	var max int
	fmt.Print("Type max fibonacci number:")
	fmt.Scan(&max)
	fmt.Printf("The sum the even fibonacci numbers below \033[32m%d\033[0m is \033[32m%d\033[0m",max,fibonnaci_even_sum(max))
}

func fibonnaci_even_sum(max int) int{
	var last_numbers [2]int = [2]int{1,2}
	var sum int = 2
	var new_fib_num int
	for {
		new_fib_num = sum_array(last_numbers)
		if new_fib_num <= max{
			last_numbers[0] = last_numbers[1]
			last_numbers[1] = new_fib_num
			if new_fib_num % 2 == 0{
				sum += new_fib_num
			}
		}else{
			return sum
		}
	}
}

func sum_array(arr [2]int) int{
	var sum int
	for _, num := range arr{
		sum += num
	}
	return sum
}