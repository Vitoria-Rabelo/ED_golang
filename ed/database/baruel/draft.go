package main
import "fmt"
func main() {
    var figAlbum, figBaruel int

    fmt.Scan(&figAlbum, &figBaruel)

    var vector []int
    for i := 0; i < figBaruel; i++ {
        var num int
        fmt.Scan(&num)
        vector = append(vector, num)
    }

    for i := 0; i < figBaruel; i++ {
        for u := 0; u < figBaruel; u++ {
           if vector[i] < vector[u] {
            vector[i], vector[u] = vector[u], vector[i]
            }
        }
    }
    
    var boolean, primeirafig bool
    for i := 0; i < figBaruel; i++ {
        for u := 0; u < figBaruel; u++ {
                if vector[i] == vector[u] && i != u && vector[u] != 0 && vector[i] != 0 {
                    if primeirafig == false && boolean == false {
                        fmt.Printf("%d", vector[i])
                    }else{
                        fmt.Printf(" %d", vector[i])
                    }
                vector[i] = 0
                boolean = true
                primeirafig = true
                break
            }
        }
    }

    if boolean == false {
        fmt.Printf("N")
    }

    fmt.Println()

    var falta bool = true
    var primeiroItem bool = true
    var aux bool 
    for i := 1; i < figAlbum + 1; i++{
        for j := 0; j < figBaruel; j++{
            if vector[j] == i{
                falta = false   
                break
            }
        }
        if falta == true {
            if primeiroItem == true {
                fmt.Printf("%d", i)
                primeiroItem = false
            } else{
                fmt.Printf(" %d", i)
            }
            aux = true
        }
        falta = true
    }
    if aux == false {
        fmt.Printf("N")
    } 
    fmt.Println()
  
}
