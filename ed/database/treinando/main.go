package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func __tostr(vet []int) string {
	if len(vet) == 0 {
		return ""
	}

	if len(vet) == 1 {
		return fmt.Sprint(vet[0])
	}

	return fmt.Sprint(vet[0]) + ", " + __tostr(vet[1:])
	
	
} 

func tostr(vet []int) string {
	
	coisa := __tostr(vet)
	return "[" + coisa + "]"

}

func __tostrrev(vet []int) string {
	if len(vet) == 0 {
		return ""
	}

	if len(vet) == 1 {
		return fmt.Sprint(vet[0])
	}

	return fmt.Sprint(vet[len(vet)-1]) + ", " + __tostrrev(vet[:len(vet)-1])
}

func tostrrev(vet []int) string {
	return "[" + __tostrrev(vet) + "]"
}


func reverse(vet []int) {
	for i, j := 0, len(vet)-1; i < j; i, j = i+1, j-1 {
		vet[i], vet[j] = vet[j], vet[i]
	}

}

func sum(vet []int) int {
	var contador int 

	for _,elem := range vet{
		contador += elem
	}

	return contador
}

func mult(vet []int) int {
	var multiplica int = 1

	for _, elem := range vet{
		multiplica *= elem
	}

	return multiplica
}

func min(vet []int) (int, int) {
	if len(vet) == 0{
		return 0 , -1
	}

	var indice int = 0
	menor := vet[0]

	for i, elem := range vet{
		if elem  < menor {
			menor = elem
			indice = i
		}
	}
	
	return menor, indice
	
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			_, index := min(vet)
			fmt.Println(index)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
