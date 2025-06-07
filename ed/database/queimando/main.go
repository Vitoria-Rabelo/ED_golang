package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

// Retorna os vizinhos (acima, abaixo, esquerda, direita)
func getNeib(p Pos) []Pos {
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c - 1},
		{p.l, p.c + 1},
	}
}

// Verifica se a posição está dentro da matriz
func inside(mat [][]rune, p Pos) bool {
	if p.l >= 0 && p.l < len(mat) && p.c >= 0 && p.c < len(mat[0]){
		return true
	}
	return false
	/* Essa função retorna uma booleana que
	   diz se a posição p está dentro da
	   matriz mat */
}

// Função para propagar a "chama" (substitui '#' por 'o')
func queimar(mat [][]rune, l, c int) {
	stack := NewStack[Pos]()

	 if inside(mat, Pos{l, c}) && mat[l][c] == '#' {
		stack.Push(Pos{l, c})
	}

	 
	
	for !stack.IsEmpty() {
		primeiraPos := *stack.Top()
		stack.Pop()
											
		if mat[primeiraPos.l][primeiraPos.c] == '#' {
			mat[primeiraPos.l][primeiraPos.c] = 'o'
			for _, vizinho := range getNeib(primeiraPos) {
				if inside(mat, vizinho) && mat[vizinho.l][vizinho.c] == '#' {
					stack.Push(vizinho)
				}
			}
		 }
	}
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como
	// queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha se possível tocar fogo nele
	// - enquanto a pilha não estiver vazia:
	//   - pegar o elemento do topo
	//   - ver se tem algum de seus vizinhos que pode ser queimado
	//   - se sim, tocar fogo nele e adicionar na pilha
	//   - se não, remover o elemento do topo da pilha
}

func main() {
	var nl, nc, l, c int
	fmt.Scan(&nl, &nc, &l, &c)

	mat := make([][]rune, nl)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for i := 0; i < nl; {
		if scanner.Scan() {
			line := scanner.Text()
			if len(line) == nc {
				mat[i] = []rune(line)
				i++
			}
		}
	}

	queimar(mat, l, c)
	showMat(mat)
}

func showMat(mat [][]rune) {
	for _, row := range mat {
		fmt.Println(string(row))
	}
}
