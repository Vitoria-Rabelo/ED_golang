package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
)

type Vector struct{
	data []int
	size int
	capacity int
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%v capacity:%v", v.size, v.capacity)
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

func (v *Vector) GetSize() int{
	return v.size
}

func (v *Vector) Capacity() int{
	return v.capacity
}

func NewVector(capacity int) * Vector {
	return &Vector{
		size: 0,
		data: make([]int, capacity),
		capacity: capacity,
	}
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity > v.capacity{
		newData := make([]int, newCapacity)

		copy(newData, v.data)
		v.data = newData
		v.capacity = newCapacity
	}

}
func (v *Vector) PushBack(value int){
	if v.size == v.capacity {
		v.Reserve(max(1, v.capacity*2))
	}

	v.data[v.size] = value
	v.size++
}

func (v *Vector) Clear() {
	 v.size = 0
}

func (v * Vector) PopBack()  error{
	if v.size == 0{
		return errors.New("vector is empty")
	}
	v.size--
	return nil
}

func (v *Vector) Insert(value int, index int) error{
	if index < 0 || index > v.size{
		return fmt.Errorf("index out of size")
	}

	if v.size == v.capacity{
		v.Reserve(max(1, v.capacity * 2))
	}

	for i := v.size; i > index; i--{
		v.data[i] = v.data[i - 1]
	}

	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector) Get(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, errors.New("index out of range")
	}
	return v.data[index], nil
}

func (v * Vector) Set(index int, value int) error{
	if index < 0 || index >= v.size {
        return errors.New("index out of range")
    }
	v.data[index] = value
	return nil
}

func (v * Vector) Erase(index int) error{
	if index < 0 || index >= v.size {
        return errors.New("index out of range")
    }

	for i := index; i < v.size-1; i++ {
        v.data[i] = v.data[i+1]
    }
	v.size--

	return nil
}

func (v * Vector) IndexOf(value int) int{
	for i := 0; i < v.size; i++{
		if value == v.data[i]{
			return i
		}
	}
	return -1
}

func (v * Vector) Contains(value int) bool{
	for i := 0; i < v.size; i++{
		if value == v.data[i]{
			return true
		}
	}
	return false
}

func (v *Vector) Reverse() {
	for i, j := 0, v.size-1; i < j; i, j = i+1, j-1 {
        v.data[i], v.data[j] = v.data[j], v.data[i]
    }
}

func (v * Vector) Sort(){
    for i := 0; i < v.size -1; i++ {
        for j := 0; j < v.size - i - 1; j++ {
            if v.data[j] > v.data[j+1] {
                v.data[j], v.data[j+1] = v.data[j+1], v.data[j]
            }
        }
    }
}

func (v *Vector) Sorted() *Vector {

    new := NewVector(v.capacity)
    new.size = v.size

    for i := 0; i < v.size; i++ {
        new.data[i] = v.data[i]
    }
    for i := 0; i < new.size-1; i++ {
        for j := 0; j < new.size-i-1; j++ {
            if new.data[j] > new.data[j+1] {
                new.data[j], new.data[j+1] = new.data[j+1], new.data[j]
            }
        }
    }
    
    return new
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(value, index)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
			 	fmt.Println("true")
			} else {
			 	fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
		 fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.Get(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			 if err != nil {
				fmt.Println(err)
			 }
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "sort":
			v.Sort()
		case "sorted":
			sorted := v.Sorted()
    		fmt.Println(sorted)
		case "reverse":
			v.Reverse()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

