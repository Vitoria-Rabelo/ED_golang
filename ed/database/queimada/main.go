package main

import (
	"bufio"
	"fmt"
	"os"
)

func tocarFogo(mat [][]rune, l, c int) {
	nl := len(mat)
	nc := len(mat[0])

	if l >= nl || c >= nc || l < 0 || c < 0{
		return
	} 

	if mat[l][c] != '#' {
		return
	}

	mat[l][c] = 'o'

	tocarFogo(mat, l-1, c)
	tocarFogo(mat,l , c-1)
	tocarFogo(mat, l+1, c)
	tocarFogo(mat,l , c+1)
	
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	mat := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		linha := []rune(scanner.Text())
		mat = append(mat, linha)
	}
	tocarFogo(mat, lfire, cfire)
	showMat(mat)
}

func showMat(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
