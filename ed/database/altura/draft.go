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

func (bt * Btree) Find(value int) *Node{
    return find(bt.Root, value)
} 

func find(node * Node, value int) *Node{
    if node == nil{
        return nil
    }

    if node.Value == value{
        return node
    }

    aux := find(node.Left, value) 
    if aux != nil{
        return aux
    }

    return find(node.Right, value)
}

func (bt * Btree) GetAltura(node * Node) int{
    if node == nil{
        return 0
    }

    leftHeight := bt.GetAltura(node.Left)
    rightHeight := bt.GetAltura(node.Right)
    if leftHeight > rightHeight {
		return leftHeight + 1
	}
    return rightHeight + 1
} 

func (bt * Btree) GetProfundidade(node * Node) int{
    return getProfundidade(bt.Root, node, 0)
}

func getProfundidade(temp, aux * Node, profundidade int) int{
    if temp == nil{
        return 0
    }

    if temp == aux{
        return profundidade + 1
    }

    aux2 := getProfundidade(temp.Left, aux, profundidade + 1)
    if aux2 != 0{
        return aux2
    }
    return getProfundidade(temp.Right, aux, profundidade + 1)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	serial := scanner.Text()
    bt := NewBTree()
    bt.Deserialize(serial)

    scanner.Scan()
	valuesStr := scanner.Text()
	values := strings.Fields(valuesStr)

    for _, valStr := range values {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Println("-1")
			continue
		}
		node := bt.Find(val)
		if node == nil {
			fmt.Println("-1")
		} else {
			altura := bt.GetAltura(node)
			profundidade := bt.GetProfundidade(node)
			fmt.Printf("%d %d\n", altura, profundidade)
		}
	}
    
}
