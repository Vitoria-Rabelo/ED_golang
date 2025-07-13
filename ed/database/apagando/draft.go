package main
import "fmt"

func main() {
   var n int
   fmt.Scan(&n)

   vet := make([]int, n)

   for i := range n{
      fmt.Scan(&vet[i])
   }



   var sair int

   fmt.Scan(&sair)

   vet_sair := make(map[int]struct{})

   for range sair{
      var elemento int
        fmt.Scan(&elemento)
        vet_sair[elemento] = struct{}{}
   }

   for _, elem := range vet {
        if _, existe := vet_sair[elem]; !existe {
            fmt.Print(elem, " ")
        }
    }
    fmt.Println()
}

