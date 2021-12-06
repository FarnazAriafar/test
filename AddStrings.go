package main

import (
	"fmt"
	"strconv"
)
func main() {
	num1, num2 := "11", "1234"
	var first uint = uint(stringToInteger1(num1))
	var second uint = uint(stringToInteger1(num2))
	sum := first + second
	var result string = integerToString1(sum)
	fmt.Println(result)
}
func stringToInteger1(s string) int {
	length := len(s)
	//a := []int{}
	var sum1 int = 0
	for i := 0; i < length; i++ {
		res, e := strconv.Atoi(string([]rune(s)[i]))
		sum1+= res
		if i != (length - 1) {
			sum1= sum1*10
		}
		if e == nil {
		}
	}
	return sum1
}
func integerToString1(a uint) string {
	var bytes []byte
	if a == 0 {
		return "0"
	}
	for a != 0 {
		bytes = append([]byte{(byte)(48 + a % 10)}, bytes...)
		a = a/ 10
	}

	var str string = string(bytes)
	return str
}