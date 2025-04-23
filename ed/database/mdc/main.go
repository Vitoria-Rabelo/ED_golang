package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0 {
		return b
	}else if b == 0{
		return a
	}else{

		return mdc(b,a % b)
	}
	return 0
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
