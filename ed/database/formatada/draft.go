package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

type BTree struct {
	Root *Node
}

func NewBTree() *BTree {
	return &BTree{
		Root: nil,
	}
}

func (bt *BTree) Deserialize(serial string) {
	scanner := bufio.NewScanner(strings.NewReader(serial))
	scanner.Split(bufio.ScanWords)
	bt.Root = deserialize(scanner)
}

func deserialize(scanner *bufio.Scanner) *Node {
	if !scanner.Scan() {
		return nil
	}
	token := scanner.Text()
	if token == "#" {
		return nil
	}
	value, err := strconv.Atoi(token)
	if err != nil {
		return nil
	}
	node := NewNode(value)
	node.Left = deserialize(scanner)
	node.Right = deserialize(scanner)
	return node
}

func (bt *BTree) PrintInOrderWithIndent() {
	printInOrderWithNulls(bt.Root, 0)
}

func printInOrderWithNulls(node *Node, depth int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		printInOrderWithNulls(node.Left, depth+1)
	} else if node.Right != nil {
	
		for i := 0; i < depth+1; i++ {
			fmt.Print("....")
		}
		fmt.Println("#")
	}
	
	for i := 0; i < depth; i++ {
		fmt.Print("....")
	}
	fmt.Println(node.Value)
    	if node.Right != nil {
		printInOrderWithNulls(node.Right, depth+1)
	} else if node.Left != nil {
		for i := 0; i < depth+1; i++ {
			fmt.Print("....")
		}
		fmt.Println("#")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	bt := NewBTree()
	bt.Deserialize(line)
	bt.PrintInOrderWithIndent()
}