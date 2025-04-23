package main
import "fmt"

func dividir(num int){
    if num == 0 {
        return
    }

    resultado := num/2
    resto := num % 2

    dividir(resultado)

    fmt.Println( resultado, resto)

}

func main() {
    var num int
    fmt.Scanf("%v", &num)

    dividir(num)
}
