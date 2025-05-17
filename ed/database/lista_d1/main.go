package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
	value int
	next *Node
}

type List struct{
	head *Node
}

func NewNode(value int, next *Node) *Node{
	return &Node{
		value: value,
		next: next,
	}
}

func NewList() *List{
	return &List{
		head: nil,
	}
}

func pushBack(node *Node, value int) *Node{
	if node == nil{
		novo := NewNode(value, nil)
		return novo
	}

	node.next = pushBack(node.next, value)
	return node 
}

func popBack (node *Node) *Node{
	if node == nil{
		return nil
	}

	if node.next == nil{
		return nil
	}

	 if node.next.next == nil {
        node.next = nil
        return node
    }

	node.next = popBack(node.next)
	return node 
}

func (l * List) PushBack(value int){
	l.head = pushBack(l.head, value)
}
func (l * List) PopBack(){
	 if l.head == nil {
        return
    }

     if l.head.next == nil {
        l.head = nil
        return
    }

	l.head = popBack(l.head)
}

func toString(node * Node) string{
	if node == nil {
		return "[]"
	}
	elementos := "["
    current := node
    first := true
    
    for current != nil {
        if !first {
            elementos += ", "
        }
        elementos += fmt.Sprint(current.value)
        current = current.next
        first = false
    }
    
    elementos += "]"
    return elementos
}

func size(node * Node) int {
	if node == nil{
		return 0
	}

	return 1 + size(node.next)
}

func (l * List) Size() int{
	return size(l.head)
}

func (l *List) PushFront(value int) {
	l.head = NewNode(value, l.head)
}

func (l * List) PopFront() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *List) Clear() {
	l.head = nil
}

func (l * List) String() string{
	return toString(l.head)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			 for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
