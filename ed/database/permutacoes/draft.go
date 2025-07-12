package main
import "fmt"
import "sort"

func permutar(s []rune, l, r int, resultados*[]string){
    if l == r{
        *resultados = append(*resultados,string(s))
    }else{
        for i := l; i <= r; i++{
            s[l], s[i] = s[i], s[l]
            permutar(s, l+1, r, resultados)
            s[l], s[i] = s[i], s[l]
        }
    }
}
func main() {
    var entrada string
    fmt.Scan(&entrada)

    caracteres := []rune(entrada)
    var resultados []string

    permutar(caracteres, 0, len(caracteres)-1, &resultados)

   sort.Strings(resultados)

    for _, perm := range resultados {
        fmt.Println(perm)
    }

}
