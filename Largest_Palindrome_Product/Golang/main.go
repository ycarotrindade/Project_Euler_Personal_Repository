package main

import (
	"fmt"
	"strconv"
)

func main(){
	var largest_palindrome int
	for i := 999; i >= 100; i--{
		for j := 999; j >= 100; j--{
			var num int = i * j
			if identify_palindrome(num){
				if num > largest_palindrome{
					largest_palindrome = num
				}
			}
		}
	}
	fmt.Print(largest_palindrome)
}

func identify_palindrome(n int) bool{
	var str_number string = strconv.Itoa(n)
	var reversed string
	for _, v := range str_number{
		reversed = string(v) + reversed
	}
	if reversed == str_number{
		return true
	}else{
		return false
	}
}