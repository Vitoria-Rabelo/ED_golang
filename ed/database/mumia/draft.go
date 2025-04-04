package main

import "fmt"

func categoria(idade int) string {
    var x string = "mumia"
    if idade < 12 {
        x = "crianca"
        return x
    }else if idade < 18 {
        x = "jovem"
        return x
    }else if idade < 65 {
        x = "adulto"
        return x
    }else if idade  < 1000 {
        x = "idoso"
    }
    return x

}
func main() {
    var nome string 
    fmt.Scan(&nome)

    var idade int
    fmt.Scan(&idade)

    var pessoa = categoria(idade)
    fmt.Print(nome, " eh ", pessoa , "\n")
}
