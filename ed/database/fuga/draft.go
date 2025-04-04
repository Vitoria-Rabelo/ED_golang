package main
import "fmt"

func main() {
    var h, p, f, d int 
    fmt.Scan(&h, &p, &f, &d)
   
    var fugiu bool

    if d == -1 {
        if h > f && h < p {
            fugiu = true
        }

        if h < f && p > f {
            fugiu = true
        }

        if h > p && p > f {
            fugiu = true
        }

    }

    if d == 1 {
        if h > f && h < p {
            fugiu = false
        }

        if h < f && p > f {
            fugiu = false
        }

        if h > p && p > f {
            fugiu = false
        }
    }

    if fugiu == true {
        fmt.Printf("S\n")
    }else{
        fmt.Printf("N\n")
    }

}
