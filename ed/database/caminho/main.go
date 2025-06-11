package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) pegarVizinhos() []Pos {
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c - 1},
		{p.l, p.c + 1},
	}
}

func procurarSaida(mat [][]rune, inicio Pos, fim Pos) {
	fila := NewQueue[Pos]()
    fila.Enqueue(inicio)

	visitados := make(map[Pos]bool)
	visitados[inicio] =true
	antes := make(map[Pos]Pos)

	for !fila.IsEmpty() {
		atual, ok := fila.Dequeue()
		if !ok {
			break
		}
		if atual == fim {
			break
		}

		for _, v := range atual.pegarVizinhos() {
			if v.l >= 0 && v.l < len(mat) && v.c >= 0 && v.c < len(mat[0]) {
				if (mat[v.l][v.c] == ' ' || mat[v.l][v.c] == 'F') && !visitados[v] {
					visitados[v] = true
					antes[v] = atual
          
					fila.Enqueue(v)   
				}
			}
		}

	}

	if _, encontrado := visitados[fim]; encontrado {
		for atual := fim; atual != inicio; atual = antes[atual]{
			mat[atual.l][atual.c] = '.'
		}
		mat[inicio.l][inicio.c] = '.'}
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
