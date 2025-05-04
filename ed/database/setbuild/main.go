package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Vector struct{
	data []int
	size int
	capacity int
}

func (v *Vector) String() string{
	elementos := "["

	for i := range v.size{
		if i != 0 {
			elementos += ", "
		}
		elementos += fmt.Sprint(v.data[i])
	}
	elementos += "]"
	return elementos
}

func NewSet(capacity int) * Vector {
	return &Vector{
		size: 0,
		data: make([]int, capacity),
		capacity: capacity,
	}
}

func (v * Vector) Insert(value int) {

	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return
		}
	}
	
	if v.size == v.capacity{
		v.Reserve(max(1, v.capacity * 2))
	}
	
	pos := v.size
	for i := 0; i < v.size; i++ {
		if v.data[i] > value {
			pos = i
			break
		}
	}

	for i := v.size; i > pos; i-- { 
		v.data[i] = v.data[i - 1]
	}
	v.data[pos] = value
	v.size++
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity > v.capacity{
		newData := make([]int, newCapacity)

		copy(newData, v.data)
		v.data = newData
		v.capacity = newCapacity
	}
}

func (v *Vector) Clear() {
	v.size = 0
}

func (v * Vector) Erase(value int) error{
	if !v.Contain(value) {
		return fmt.Errorf("value not found")
	}

	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			for j := i; j < v.size-1; j++ {
				v.data[j] = v.data[j+1]
			}
			v.size--
			return nil
		}
	}
	return nil
	
}

func (v * Vector) Contain(value int) bool{
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return true
		}
	}
	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value,_ := strconv.Atoi(parts[1])
			 v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("fail: invalid value", part)
					continue
				}
			v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if err := v.Erase(value); err != nil {
				fmt.Println(err) 
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contain(value){
				fmt.Println("true")
		   } else {
				fmt.Println("false")
		   }
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
