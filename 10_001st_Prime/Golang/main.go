package main

import "fmt"

func main() {
	var res []int = sieve_of_eratosthenes(100000000)
	fmt.Print(res[10000])
}

func sieve_of_eratosthenes(n int) []int{
	sieve := make([]bool,n+1)
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