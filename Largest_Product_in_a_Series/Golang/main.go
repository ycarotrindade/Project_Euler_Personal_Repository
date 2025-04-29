package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var str_number string
	file, _ := os.Open("../number.txt")
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan(){
		str_number += s.Text()
	}
	var str_len int = utf8.RuneCountInString(str_number)
	var max_adjacents int = 13
	var actual_numbers []string
	var product int = 1
	var largest_product int
	for i := range str_number{
		product = 1
		if i + max_adjacents < str_len{
			actual_numbers = strings.Split(str_number[i:i+max_adjacents],"")
			for i := range actual_numbers{
				value, _ := strconv.Atoi(actual_numbers[i])
				product *= value
			}
			if product > largest_product{
				largest_product = product
			}
		}else{
			break
		}
	}
	fmt.Println(largest_product)
}