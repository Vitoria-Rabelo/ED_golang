package main

import "fmt"

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}

func min(a, b int)int{
    if a < b {
        return a
    }
    return b
}

func backtracking(seq []byte, pos int, d byte, L int) bool {
    for r := max(0, pos-L); r < pos; r++ {
        if seq[r] == d {
            return false
        }
    }
    for z := pos + 1; z <= min(len(seq)-1, pos+L); z++ {
        if seq[z] == d {
            return false
        }
    }
    return true
}

func distancia(seq []byte, L int) bool{
    for i := 0; i < len(seq); i++ {
        if seq[i] == '.' {
            for j := 0; j <= L; j++ {
                d := byte(j + '0')
                if backtracking(seq, i, d, L) {
                    seq[i] = d
                    if distancia(seq, L) {
                        return true
                    }
                    seq[i] = '.'
                }
            }
            return false
        }
    }
    return true
}

func main() {
    var sequencia string
    fmt.Scan(&sequencia)

    var l int
    fmt.Scan(&l)
    seq := []byte(sequencia)
    
    if distancia(seq, l) {
        fmt.Println(string(seq))
    }
    
    slice := make([]int, 0, l)
    for i := 0; i < l; i++ {
        slice = append(slice, i)
    }
}