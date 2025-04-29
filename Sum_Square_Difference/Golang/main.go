package main

import (
	"fmt"
	"math"
)

func main(){
	var res_1 int = sum_of_arithmetic_series(100,1,1)
	var res_2 int = sum_of_square_arithmetic_series(100,1,1)
	res_1 = int(math.Pow(float64(res_1),2))
	var diff int = res_1 - res_2
	fmt.Println(diff)
}

func sum_of_arithmetic_series(n int, a1 int, r int) int{
	var an int
	var sn int
	an = a1 + (n - 1 ) * r
	sn = (n * (a1 + an)) / 2
	return sn
}

func sum_of_square_arithmetic_series(n int, a1 int, r int) int{
	var an int
	var sn int
	an = a1 + (n - 1 ) * r
	for i := a1; i <= an; i += r{
		sn += i * i
	}
	return sn
}