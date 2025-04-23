package main
import "fmt"
func main() {
    var tam int
    fmt.Scan(&tam)

    entrada := make([]string, tam)

    for i := 0; i < tam; i++{
        fmt.Scan(&entrada[i])
    }

    var tam2 int
    fmt.Scan(&tam2)

    consulta := make([]string, tam2)

    for i := 0; i < tam2; i++{
        fmt.Scan(&consulta[i])
    }

    for i, elem := range consulta {
        if i != 0 {
            fmt.Printf(" ")
        }
        contagem := 0
        for _, value := range entrada {
            if elem == value {
                contagem++
            }
        }

        fmt.Printf("%v", contagem)
        
    }
    fmt.Println()
}
