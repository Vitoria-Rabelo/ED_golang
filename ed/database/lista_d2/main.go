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
	prev *Node
}

type LList struct{
	root * Node
	size int
}

func NewNode(value int, next *Node, prev * Node) * Node{
	return &Node{
		value: value,
		next: next,
		prev: prev,
	}
}

func NewLList() *LList{
		list := &LList{
			root: NewNode(0, nil, nil),
		}
		list.root.next = list.root
		list.root.prev = list.root
		return list

}

func (l *LList) PushBack(value int) {
	newNode := NewNode(value, l.root, l.root.prev)
	l.root.prev.next = newNode
	l.root.prev = newNode
}

func (l *LList) PushFront(value int) {
	newNode := NewNode(value, l.root.next, l.root)
	l.root.next.prev = newNode
	l.root.next = newNode
}

func (l *LList) PopBack() {
	if l.root.prev == l.root {
		return 
	}
	last := l.root.prev
	last.prev.next = l.root
	l.root.prev = last.prev
}

func (l *LList) Replace(oldValue, newValue int) bool {
    for node := l.root.next; node != l.root; node = node.next {
        if node.value == oldValue {
            node.value = newValue
            return true 
        }
    }
    return false
}

func (l *LList) Search(value int) *Node {
	for node := l.root.next; node != l.root; node = node.next {
		if node.value == value {
			return node
		}
	}
	return nil
}


func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) Walk() {
	fmt.Print("[ ")
	for node := l.root.next; node != l.root; node = node.next {
		fmt.Printf("%v ", node.value)
	}
	fmt.Println("]")
	fmt.Print("[ ")

	for node := l.root.prev; node != l.root; node = node.prev {
		fmt.Printf("%v ", node.value)
	}
	fmt.Println("]")
}

func (l *LList) Size() int {
  return l.size
}

func (l *LList) Insert(existing, newValue int) bool {
    node := l.Search(existing)

    if node == nil {
        return false 
    }

    newNode := &Node{
        value: newValue,
        next: node,      
        prev: node.prev, 
    }
    
    node.prev.next = newNode 
    node.prev = newNode      
    return true
}

func (l *LList) Remove(node *Node) bool {
	if node == nil || node == l.root {
		return false
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
	return true
}

func toString(node *Node) string {
    if node == nil || node.next == node {
        return "[]"
    }
    
    var elementos []string
    current := node.next 
    
    for current != node {
        elementos = append(elementos, fmt.Sprint(current.value))
        current = current.next
    }
    
    return "[" + strings.Join(elementos, ", ") + "]"
}

func (l * LList) String() string{
	return toString(l.root)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

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
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			ll.Walk()
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
				node.value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(oldvalue, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			 oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Remove(node)
			} else {
		 	fmt.Println("fail: not found")
			 }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
