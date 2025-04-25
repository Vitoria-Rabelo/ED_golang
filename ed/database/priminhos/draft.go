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
func primos(num int) []int{
    primos := make([]int, 0, num)

    num2 := 2

    for len(primos) < num {
        if eh_primo(num2, 2) {
            primos = append(primos, num2)
        }
        num2++
    }
    return primos
}

func main() {
    var num int
    fmt.Scan(&num)

    lista := primos(num)
    fmt.Print("[")
    for i, elem := range lista {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(elem)
    }
    fmt.Println("]")
    
}
