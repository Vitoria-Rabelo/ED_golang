package main

import "fmt"

func backtracking(vet []int, index, soma, alvo int) bool {
	if soma == alvo{
		return true
	}

	if soma > alvo{
		return false
	}

	if index >= len(vet){
	return false
	}
	
	if  backtracking(vet, index+1, soma, alvo) || backtracking(vet, index+1, soma+vet[index], alvo){
		return true
	}

	return false
}

func main() {
	var qtd, soma int
	fmt.Scan(&qtd, &soma)
	vet := make([]int, qtd)
	for i := range qtd {
		fmt.Scan(&vet[i])
	}
	if backtracking(vet, 0, 0, soma) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
