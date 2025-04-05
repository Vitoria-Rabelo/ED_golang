package main
import (
    "fmt"
    "math"
)

func main() {
    var competidores int
    fmt.Scan(&competidores)

    if competidores < 1 || competidores > 1000 {
        return
    }

    var a, b, indice int
    menorPontuacao := 101
    indice = -1

    for i := 0; i < competidores; i++ {
        fmt.Scan(&a, &b)

        if a < 10 || b < 10 {
            continue 
        }

        diff := int(math.Abs(float64(a - b)))
        if diff < menorPontuacao {
            menorPontuacao = diff
            indice = i
        }
    }
        fmt.Println(indice)
    
}

