package main

import "fmt"

func main() {
	var is_divisible_by_all bool
	var number int = 21
	for !is_divisible_by_all{
		is_divisible_by_all = true
		for i:=2; i <= 20; i++{
			if !verify_factor(number, i){
				is_divisible_by_all = false
				number++
				break
			}
		}
	}
	fmt.Print(number)
}

func verify_factor(n1 int, n2 int) bool{
	if n1 % n2 == 0{
		return true
	}else{
		return false
	}
}