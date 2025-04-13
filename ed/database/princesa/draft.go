package main

import "fmt"

func main() {
    var N, E int
    fmt.Scan(&N, &E)

    pessoas := make([]bool, N)
    for i := range pessoas {
        pessoas[i] = true
    }

    pos := E - 1
    vivas := N

    for vivas > 1 {
        fmt.Print("[")
        for i, viva := range pessoas {
            if viva {
                if i == pos {
                    fmt.Printf(" %d>", i+1)
                } else {
                    fmt.Printf(" %d", i+1)
                }
            }
        }
        fmt.Println(" ]")

        vitima := (pos + 1) % N
        for !pessoas[vitima] {
            vitima = (vitima + 1) % N
        }

        pessoas[vitima] = false
        vivas--

        pos = (vitima + 1) % N
        for !pessoas[pos] {
            pos = (pos + 1) % N
        }
    }

    fmt.Print("[")
    for i, viva := range pessoas {
        if viva {
            fmt.Printf(" %d>", i+1)
        }
    }
    fmt.Println(" ]")
}