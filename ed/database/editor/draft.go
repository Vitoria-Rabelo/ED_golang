package main

import "fmt"


func main() {
	var entrada string
    fmt.Scan(&entrada)

    slice := make([]rune, 0)
    p := 0

    for _, elem := range entrada{
        switch {
            case ( elem >= 'a' && elem <= 'z') || elem == '-' :
                slice = append(slice, 0) 
                copy(slice[p+1:], slice[p:])
                slice[p] = elem
                p++

            case elem == 'R':
                slice = append(slice, 0) 
                copy(slice[p+1:], slice[p:])
                slice[p] = '\n'
                p++
            
            case elem == 'B':
                if p > 0 {
                    slice = append(slice[:p-1], slice[p:]...)
			        p--
                }
            case elem == '<':
                if p > 0{
                    p--
                }
            
            case elem == 'D':
                    if p < len(slice){
                    slice = append(slice[:p], slice[p+1:]...)
                    
                    }
            case elem == '>':
                if p < len(slice){
                    p++
                }

        }
    }
    fmt.Println(string(slice[:p]) + "|" + string(slice[p:]))
}