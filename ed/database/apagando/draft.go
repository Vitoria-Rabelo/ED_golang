package main
import "fmt"

func main() {
    var tam int
    fmt.Scan(&tam)

    slice := make([]int, tam)

    for i := range tam{
        fmt.Scan(&slice[i])
    }

    fmt.Println(slice)

    var sai int
    fmt.Scan(&sai)

    slice_sair := make(map[int]struct{})

    for range sai{
        var elemento int
        fmt.Scan(&elemento)
        slice_sair[elemento] = struct{}{}
    }
    fmt.Println(slice_sair)

    imprime := make([]int, 0, tam)
    for _, elem := range slice{
        _, esta := slice_sair[elem]
        if !esta{
            imprime = append(imprime, elem)
        }
    }
    fmt.Println(slice)
}
