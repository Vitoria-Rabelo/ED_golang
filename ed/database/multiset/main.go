package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct{
	data []int
	size int
	capacity int
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func NewMultiSet(capacity int) * Set {
	return &Set{
		size: 0,
		data: make([]int, capacity),
		capacity: capacity,
	}
}

func (ms *Set) String() string{
	elementos := "["

	for i := range ms.size{
		if i != 0 {
			elementos += ", "
		}
		elementos += fmt.Sprint(ms.data[i])
	}
	elementos += "]"
	return elementos
}

func (ms *Set) Reserve(newCapacity int) {
	if newCapacity > ms.capacity{
		newData := make([]int, newCapacity)

		copy(newData, ms.data)
		ms.data = newData
		ms.capacity = newCapacity
	}
}

func (ms * Set) Insert(value int) {
	
	if ms.size == ms.capacity{
		ms.Reserve(max(1, ms.capacity * 2))
	}
	
	pos := ms.size
	for i := 0; i < ms.size; i++ {
		if ms.data[i] > value {
			pos = i
			break
		}
	}

	for i := ms.size; i > pos; i-- { 
		ms.data[i] = ms.data[i - 1]
	}
	ms.data[pos] = value
	ms.size++
}

func (ms * Set) Contain(value int) bool{
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			return true
		}
	}
	return false
}

func (ms * Set) Erase(value int) error{
	if !ms.Contain(value) {
		return fmt.Errorf("value not found")
	}

	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			for j := i; j < ms.size-1; j++ {
				ms.data[j] = ms.data[j+1]
			}
			ms.size--
			return nil
		}
	}
	return nil
}

func (ms * Set) Count(value int) int{
	var elem int = 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			elem++
		}
	}
	return elem
}

func (ms * Set) Clear(){
	ms.size = 0
}

func (ms * Set) Unique() int{

	if ms.size == 0 {
        return 0
    }

	var count int = 1
	for i := 1; i < ms.size; i++ {
        if ms.data[i] != ms.data[i-1] {
            count++
        }
    }
    return count
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "magic":
			// value, _ := strconv.Atoi(args[1])
		case "insert":
			for _, args := range args[1:] {
				value, err := strconv.Atoi(args)
				if err != nil {
					fmt.Println("fail: invalid value", args)
					continue
				}
			ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if err := ms.Erase(value); err != nil {
				fmt.Println(err) 
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contain(value){
				fmt.Println("true")
		   } else {
				fmt.Println("false")
		   }
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
