package main

import (
	"fmt"
	"strconv"
	"sync"
)
func main() {
	//fmt.Println(addStrings("11", "1234"))
	num1, num2 := "11", "1234"
	var sum int
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	go stringToInteger(&sum, &waitgroup, num1)
	go stringToInteger(&sum, &waitgroup, num2)
	waitgroup.Wait()
	//fmt.Println("here is after wait", "num1 is:" , )
	fmt.Println("resulted string is:", integerToString(sum))
	//return integerToString(sum)
}
func stringToInteger(sum *int, waitgroup *sync.WaitGroup, s string) {
	fmt.Println("Hello World")
	fmt.Println("sum is:",&sum)
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
	fmt.Println("Finished Execution")
	waitgroup.Done()
	*sum = sum1 + *sum
}
func integerToString(a int) string {
	var bytes []byte
	if a == 0 {
		return "0"
	}
	for a != 0 {
		bytes = append([]byte{(byte)(48 + a % 10)}, bytes...)
		a = a/ 10
	}
	return string(bytes)
}