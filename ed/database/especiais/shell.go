package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	var contador = make(map[int]int)

	for _, elem := range vet {
		if elem < 0 {
			elem = -elem
		}

		contador[elem]++
	}

	var pares []Pair

	for i, count := range contador {
		pares = append(pares, Pair{One: i, Two: count})
	}

	for i := 0; i < len(pares); i++{
		for j := i + 1; j < len(pares); j++ {
			if pares[i].One > pares[j].One {
				pares[i], pares[j] = pares[j], pares[i]
			}
		}
	}
	return pares
}

func teams(vet []int) []Pair {
	var resultado []Pair

	if len(vet) == 0 {
		return resultado
	}
	atual := vet[0]
	cont := 1
	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			cont++
		} else {
			resultado = append(resultado, Pair{One: atual, Two: cont})
			atual = vet[i]
			cont = 1
		}
	}
	resultado = append(resultado, Pair{One: atual, Two: cont})

	return resultado
}


func mnext(vet []int) []int {
	slice := make([]int , 0, len(vet))
 
	for i := 0; i < len(vet); i++ {
		elem := 0
		if  vet[i] > 0{
			if (i > 0 && vet[i-1] < 0) || (i < len(vet)-1 && vet[i+1] < 0){
				elem = 1
			}
		}else {
			elem = 0
		}
		slice = append(slice, elem)
	}
	return slice
}

func alone(vet []int) []int {
	slice := make([]int, 0, len(vet))
	for i := 0; i < len(vet); i++ {
		elem := 0

		if vet[i] > 0 {
			if len(vet) == 1 {
				elem = 1
			} else if i == 0 {
				if vet[i+1] > 0 {
					elem = 1
				}
			} else if i == len(vet)-1 {
				if vet[i-1] > 0 {
					elem = 1
				}
			} else {
				if vet[i-1] > 0 && vet[i+1] > 0 {
					elem = 1
				}
			}
		}
		slice = append(slice, elem)
	}
	return slice
}

func couple(vet []int) int {
	count := 0 

	for i := 0; i < len(vet); i++ {
		if vet[i] == 0 {
			continue
		}
		if vet[i]  != 0{
			for j := i + 1; j < len(vet); j++ {
				if vet[j] == -vet[i]{
					count++
					vet[j] = 0
					vet[i] = 0
					break
				}
			}
		}
	}
		return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if len(vet) < pos + len(seq){
		return false
	}

	for i := 0; i < len(seq); i++{
		if vet[pos + i] != seq[i]{
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	for i := 0; i <= len(vet)-len(seq); i++ {
        if vet[i] == seq[0] && hasSubseq(vet, seq, i) {
            return i
        }
    }
    return -1
}

func erase(vet []int, posList []int) []int {
	indicesRemover := make(map[int]bool)

	for _, pos := range posList {
		if pos >= 0 && pos < len(vet) {
			indicesRemover[pos] = true
		}
	}
	slice := make([]int, 0, len(vet) - len(indicesRemover))

	for i, valor := range vet{
		if !indicesRemover[i]{
			slice = append(slice, valor)
		}
	}
	return slice
}

func clear(vet []int, value int) []int {
	slice := make([]int, 0, len(vet))

	for _, elem := range vet {
		if elem != value {
			slice = append(slice, elem)
		}
	}
	return slice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
