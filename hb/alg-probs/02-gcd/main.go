package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Scan(&num1, &num2)

	result := gcd(num1, num2)
	fmt.Println(result)

	result = gcdRecursive(num1, num2)
	fmt.Println(result)

}

func gcd(num1 int, num2 int) int {
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	return num1
}

func gcdRecursive(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	}

	return gcdRecursive(num2, num1 % num2)
}