package main

import "fmt"

func main() {
	fmt.Println(addStrings("401716832807512840963","167141802233061013023557397451289113296441069"))
}
func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		return addStrings(num2, num1)
	}
	var result []byte
	var carry int
	for i := 0; i < len(num2); i++ { //  len(num2) >= len(num1
		idx1, idx2 := len(num1)-1-i, len(num2)-i-1
		value1, value2 := 0, 0
		if idx1 >= 0 {
			value1 = int(num1[idx1] - '0')
		}
		value2 = int(num2[idx2] - '0')
		sum := carry + value1 + value2
		carry = sum / 10
		result = append(result, '0'+byte(sum%10))
	}
	if carry > 0 {
		result = append(result, '1')
	}

	for i := 0; i < (len(result)+1)/2; i++ { // reverse the byte array
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return string(result)
}
