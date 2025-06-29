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

type Btree struct {
	Root *Node
}

func NewBTree() *Btree {
	return &Btree{
		Root: nil,
	}
}

func (bt *Btree) Deserialize(serial string) {
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

func (bt * Btree) Soma() int{
	return soma(bt.Root)
}

func soma(node * Node) int{
	if node == nil{
		return 0
	}
	aux := node.Value + soma(node.Right) +soma(node.Left)
	return aux
}

func (bt * Btree) Min() int{
	if bt.Root == nil{
		return 0
	}
	return min(bt.Root, bt.Root.Value)
}

func min(node * Node, minimo int) int {
    if node == nil {
        return minimo
    }
    
    if minimo > node.Value{
        minimo = node.Value
    }

    minimo = min(node.Right, minimo)
    minimo = min(node.Left, minimo)
    return minimo
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	serial := scanner.Text()
    bt := NewBTree()
	bt.Deserialize(serial)
	fmt.Printf("%d %d\n", bt.Soma(), bt.Min())
}