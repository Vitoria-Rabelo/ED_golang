package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func getMen(vet []int) []int {
	 var homem []int
	 for _, elem := range vet {
		 if elem > 0 {
			homem = append(homem, elem)
		 }
	 }
	 return homem
}

func getCalmWomen(vet []int) []int {
	var mulher []int 
	for _, elem := range vet {
		if elem < 0 && elem > -10{
			mulher = append(mulher, elem)
		}
	}
	return mulher
}


func sortVet(vet []int) []int {
	newsort := make([]int, len(vet))
	copy(newsort, vet)

	sort.Slice(newsort, func(i, j int) bool {
		return newsort[i] < newsort[j]
	})
	return newsort
}

func sortStress(vet []int) []int {
	newsort := make([]int, len(vet))
	copy(newsort, vet)

	sort.Slice(newsort, func(i, j int) bool {
		if newsort[j] < 0 && newsort[i] < 0 {
			return -newsort[i] < -newsort[j]
		}else if newsort[j] < 0 {
			return newsort[i] < -newsort[j]
		} else if newsort[i] < 0 {
			return -newsort[i] < newsort[j]
		}else{
			return newsort[i] < newsort[j]
		}
	})
	return newsort
}

func reverse(vet []int) []int {
	reverte := make([]int, 0 , len(vet))

	for i := len(vet) - 1; i >= 0; i--{
		reverte = append(reverte, vet[i])
	}
	return reverte
}

func reverseInplace(vet []int) {
	_ = vet
}

func unique(vet []int) []int {
    var unico []int
    
    contains := func(slice []int, valor int) bool {
        for _, elem := range slice {
            if elem == valor {
                return true
            }
        }
        return false
    }
    
    for _, elem := range vet {
        if !contains(unico, elem) {
            unico = append(unico, elem)
        }
    }
    return unico
}

func repeated(vet []int) []int {
    contador := make(map[int]int)
    numerosUnicos := make(map[int]bool)
    var repetidos []int

    for _, num := range vet {
        contador[num]++
    }

    for num, count := range contador {
        if count > 1 {
            numerosUnicos[num] = true
        }
    }

    var numeros []int
    for num := range numerosUnicos {
        numeros = append(numeros, num)
    }

    sort.Ints(numeros)

    for _, num := range numeros {
        for i := 0; i < contador[num]-1; i++ {
            repetidos = append(repetidos, num)
        }
    }

    return repetidos
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			printVec(reverse(str2vet(args[1])))
		case "reverse_inplace":
			vet := str2vet(args[1])
			reverseInplace(vet)
			printVec(vet)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

