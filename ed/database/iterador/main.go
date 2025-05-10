package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterator struct {
	Iterator
}

type CyclicIterator struct {
	Iterator
	max     int
	counter int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: 0}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *Iterator) Next() int {
	if !i.HasNext() {
		panic("No more elements")
	}
	value := i.data[i.index]
	i.index++
	return value
}

func (l *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{
		Iterator: Iterator{
			data:  l.data,
			index: len(l.data) - 1,
		},
	}
}

func (ri *ReverseIterator) HasNext() bool {
	return ri.index >= 0
}

func (ri *ReverseIterator) Next() int {
	if !ri.HasNext() {
		panic("No more elements")
	}
	value := ri.data[ri.index]
	ri.index--
	return value
}

func (l *MyList) CyclicIterator(max int) *CyclicIterator {
	return &CyclicIterator{
		Iterator: Iterator{
			data:  l.data,
			index: 0,
		},
		max: max,
	}
}

func (ci *CyclicIterator) HasNext() bool {
	return ci.counter < ci.max
}

func (ci *CyclicIterator) Next() int {
	if len(ci.data) == 0 {
		panic("Empty list")
	}
	if !ci.HasNext() {
		panic("No more elements")
	}
	value := ci.data[ci.index]
	ci.index = (ci.index + 1) % len(ci.data)
	ci.counter++
	return value
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			slice := make([]int, len(args)-1)
			for i, value := range args[1:] {
				num, _ := strconv.Atoi(value)
				slice[i] = num
			}
			mylist = NewMyList(slice)
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Print(it.Next(), " ")
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Print(it.Next(), " ")
			}
			fmt.Println("]")
		case "cyclic":
			if len(args) < 2 {
				fmt.Println("Usage: cyclic <max_iterations>")
				continue
			}
			max, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid number of iterations")
				continue
			}
			fmt.Print("[ ")
			for it := mylist.CyclicIterator(max); it.HasNext(); {
				fmt.Print(it.Next(), " ")
			}
			fmt.Println("]")
		}
	}
}