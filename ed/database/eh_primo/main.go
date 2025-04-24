package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x == 2 || div * div > x && x != 1{
		return true
	}
	if x <= 1 || x % div == 0{
		return false
	}

	return eh_primo(x, div + 1);
}

func main() {
	var x int
	fmt.Scan(&x)

	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
