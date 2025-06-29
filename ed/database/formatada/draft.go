package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct{
	Value int
	Left *Node
	Right *Node
}

type Tree struct{
	Root *Node
}

func NewNode(value int) *Node{
	return &Node{
		Value: value,
		Left: nil,
		Right: nil,
	}
}

func NewTree() *Tree{
	return &Tree{
		Root: nil,
	}
}

func (tree *Tree) Deserialize(serial string) {
	scanner := bufio.NewScanner(strings.NewReader(serial))
	scanner.Split(bufio.ScanWords)
	tree.Root = deserialize(scanner)
}

func deserialize(scanner *bufio.Scanner) *Node{
	if !scanner.Scan() {
		return nil
	}

	token := scanner.Text()
	if token == "#" {
		return nil
	}
	value, error := strconv.Atoi(token)
	if error != nil {
		return nil
	}
	n := NewNode(value)
	n.Left = deserialize(scanner)
	n.Right = deserialize(scanner)
	return n

}

func (tree * Tree) Show(){
	show(tree.Root, 0)
}

func show(node * Node, d int){
	if node == nil{
		return
	}

	if node.Left != nil{
		show(node.Left, d + 1)
	} else if node.Right != nil{
		for i := 0; i < d + 1; i++{
			fmt.Print("....")
		}
		fmt.Println("#")
	}

	for i := 0; i < d ; i++ {
		fmt.Print("....")
	}
	fmt.Println(node.Value)
    	if node.Right != nil {
			show(node.Right, d + 1)
		}else if node.Left != nil {
			for i := 0; i < d+1; i++ {
				fmt.Print("....")
			}
			fmt.Println("#")
		}

}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	tree := NewTree()
	tree.Deserialize(input)
	tree.Show()
}