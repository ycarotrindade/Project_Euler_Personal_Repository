package main

import (
	"fmt"
)

func main() {
	var max_number int
	fmt.Print("Type max number:")
	fmt.Scan(&max_number)
	var primes []int = sieve_of_eratosthenes(10000)
	var aux int = 0
	for max_number > 1{
		var prime int = primes[aux]
		if verify_factor(max_number, prime){
			max_number /= prime
			}else{
				aux++
			}
			if aux >= len(primes){
				break
			}
	}
	fmt.Println(max_number)
	fmt.Println(primes[aux])
}

func verify_factor(n1 int, n2 int) bool{
	if n1 % n2 == 0{
		return true
	}else{
		return false
	}
}

func sieve_of_eratosthenes(n int) []int{
	sieve := make([]bool, n+1)
	var primes []int
	for i := range sieve{
		sieve[i] = true
	}

	for p := 2; p * p <= n; p++{
		if sieve[p]{
			for i := p * p; i <= n; i += p{
				sieve[i] = false
			}
		}
	}
	for i, v := range sieve{
		if v{
			primes = append(primes, i)
		}
	}

	return primes[2:]
}