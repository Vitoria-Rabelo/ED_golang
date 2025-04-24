package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x == 2 || div * div > x && x != 1{
		return true
	}
	if x <= 1 || x % div == 0{
		return false
	}

	return eh_primo(x, div + 1);
}

func enesimo(n int, y int) int{
    if n == 0{
        return y - 1
    }
    if eh_primo(y ,2){
        return enesimo(n - 1, y + 1)
    }
    return enesimo(n, y + 1)
}
func main() {
	var x int
	fmt.Scan(&x)
    
    resultado := enesimo(x, 2)
    fmt.Println(resultado)

}
