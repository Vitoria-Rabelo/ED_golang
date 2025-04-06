package main
import "fmt"
func main() {
    var tam int
    fmt.Scan(&tam)

    var vector []int

    for i := 0; i < tam; i++{
        var animal int
        fmt.Scan(&animal)
        vector = append(vector, animal)
    }

    var qtdCasais int

    for x := 0; x < tam; x++{
        for y := 1; y < tam; y++{
            if vector[x] != vector[y] && vector[x] == (vector[y] * -1){
                qtdCasais++
                vector[x] = 0
                vector[y] = 0
                break
            }
        }
    }

    fmt.Printf("%d\n", qtdCasais)
}
