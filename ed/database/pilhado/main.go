package main

import (
	"bufio"
	"fmt"
	"os"
)

// Pos representa uma posição (linha, coluna) na matriz.
// Equivalente à struct Pos em C++
type Pos struct {
	l int
	c int
}

func (p Pos) pegarVizinhos() []Pos {
 // Retorna as posições vizinhas (cima, baixo, esquerda, direita)
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c - 1},
		{p.l, p.c + 1},
	}
}

// procurarSaida é uma função placeholder. Você precisará implementar a lógica
// equivalente à sua função procurar_saida em C++.
// Assumimos que ela modifica a matriz in-place.
func procurarSaida(mat [][]rune, inicio Pos, fim Pos) {
	caminho := NewStack[Pos](len(mat) * len(mat[0]))
	caminho.Push(inicio)
	beco := NewStack[Pos](len(mat) * len(mat[0]))
	visitados := make([][]bool, len(mat))
	for i := range visitados {
		visitados[i] = make([]bool, len(mat[0]))
	}

	for !caminho.IsEmpty(){
		atual, err := caminho.Top()
		
		if err != nil {
			break
		}

		visitados[atual.l][atual.c] = true

		if atual == fim{
			break
		}
		var validos []Pos

		for _, v := range atual.pegarVizinhos() {
			if v.l >= 0 && v.l < len(mat) && v.c >= 0 && v.c < len(mat[0]) {
				if (mat[v.l][v.c] == ' ' || mat[v.l][v.c] == 'F') && !visitados[v.l][v.c] {
					validos = append(validos, v)
				}
			}
		}
			if len(validos) > 0{
				caminho.Push(validos[0])
			} else {
				beco.Push(atual)
				caminho.Pop()
			}
		
		for _, pos := range caminho.data {
			mat[pos.l][pos.c] = '.'
		}
		
		for _, pos := range beco.data {
			if mat[pos.l][pos.c] == '.' {
				mat[pos.l][pos.c] = ' '
			}
		}

	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	procurarSaida(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
