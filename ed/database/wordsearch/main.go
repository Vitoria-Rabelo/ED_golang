package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	if len(word) == 0{
		return true
	}

	l := len(grid)
	c := len(grid[0])

	if l == 0{
		return false
	}

	visitado := make([][]bool, l)

	for i := range visitado {
		visitado[i] = make([]bool, c)
	}

	for i := 0; i < l; i++{
		for j := 0; j < c; j++{
			if busca(grid, word, visitado,i, j, 0){
				return true
			}
			
		}
	}
	return false
}

func busca(grid [][]byte, word string, visitado [][]bool, i, j, aux int) bool{
	if aux == len(word){
		return true
	}

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != word[aux] || visitado[i][j]{
		return false
	} 

	visitado[i][j] = true

	if busca(grid, word, visitado, i+1, j, aux+1) || busca(grid, word, visitado, i-1, j, aux+1) || busca(grid, word, visitado, i, j + 1, aux+1) || busca(grid, word, visitado, i, j - 1, aux+1){
		return true
	}

	visitado[i][j] = false
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
