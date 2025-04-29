package main

import (
	"fmt"
	"math"
)

func main() {
	var m int = 1
	var a, b, c, sum, prod int
	var is_found bool
	for !is_found{
		for n := range m{
			a, b, c = generate_pythagorean(m, n)
			sum = a + b + c
			if sum == 1000{
				prod = a * b * c
				fmt.Println(prod)
				is_found = true
				break
			}
		}
		m++
	}
}

func generate_pythagorean(n1 int,n2 int) (int,int,int){
	var a int = int(math.Pow(float64(n1),2.0) - math.Pow(float64(n2),2.0))
	var b int = 2 * n1 * n2
	var c int = int(math.Pow(float64(n1),2.0) + math.Pow(float64(n2),2.0))
	return a, b, c
}