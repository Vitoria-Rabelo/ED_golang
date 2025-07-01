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

func (bt * Btree) Find(node * Node, value int) string{
	if node == nil{
		return "!"
	}

	if node.Value == value{
		return "x"
	}

	leftP := bt.Find(node.Left, value)
	if leftP != "!"{
		return "l" + leftP
	}

	rightP := bt.Find(node.Right, value)
	if rightP != "!"{
		return "r" + rightP
	}

	return "!"
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	bt := NewBTree()
	bt.Deserialize(line)

	valueStr, _ := reader.ReadString('\n')
	valueStr = strings.TrimSpace(valueStr)
	value, _ := strconv.Atoi(valueStr)

	path := bt.Find(bt.Root, value)
	fmt.Println(path)
}
