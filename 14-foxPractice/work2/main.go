package main

import (
	"errors"
	"fmt"
)

func Add(i, j int) (int, error) {
	if i < 0 {
		return -1, errors.New("Error: i不可小於0")
	}
	if j < 0 {
		return -1, errors.New("Error: j不可小於0")
	}
	return i + j, nil
}
func Sum(i int) (int, error) {
	sum := 0
	if i < 0 {
		return -1, errors.New("Error: i不可小於0")
	}
	for index := 0; index <= i; index++ {
		sum += index
	}
	return sum, nil
}
func main() {
	fmt.Println(Add(100, -1))
	fmt.Println(Add(100, 100))
	fmt.Println(Sum(100))
	fmt.Println(Sum(-100))
}
