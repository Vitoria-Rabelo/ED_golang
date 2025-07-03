
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
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
	n := NewNode(value)
	n.left = deserialize(scanner)
	n.right = deserialize(scanner)
	return n
}

func show(n *Node) {
	if n == nil {
		return
	}

	show(n.left)
	fmt.Printf("%d ", n.value)
	show(n.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	bt := NewTree()
	scanner = bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	bt.root = deserialize(scanner)

	fmt.Print("[ ")
	show(bt.root)
	fmt.Println("]")
}