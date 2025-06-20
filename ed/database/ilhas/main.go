package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	nl := len(grid)
	nc := len(grid[0])
	
	if len(grid) == 0{
		return 0
	}
	
	cont := 0
	for i := 0; i < nl; i++{
		for j := 0; j < nc; j++{
			if grid[i][j] == '1'{
				cont++
				dfs(grid, i, j, nl, nc)
			}
		}
	}
	return cont
	
}

func dfs(grid[][] byte, i , j , nl, nc int) {
	if i >= nl || j >= nc || i < 0 || j < 0 || grid[i][j] == '0'{
		return 
	}
	grid[i][j] = '0'
	dfs(grid, i + 1,j,  nl, nc)
	dfs(grid, i, j + 1, nl, nc )
	dfs(grid, i - 1, j, nl, nc)
	dfs(grid, i, j- 1, nl, nc)

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
