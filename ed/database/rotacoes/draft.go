package main
import "fmt"
func main() {
    var tam, rot int
    fmt.Scan(&tam, &rot)

    vector := make([]int ,tam)

    for i := 0; i < tam; i++{
        fmt.Scan(&vector[i])
    }

    newRot := rot % tam
    fmt.Printf("[")
    for i, _:= range vector {
            fmt.Printf(" ")
        indice := (i + tam - newRot) % tam
        fmt.Printf("%v", vector[indice])
        
    }
    fmt.Println(" ]")
}
