package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	var x int = 0

	for i, elem := range slice{
		if elem == value{
			x = i
		}
	}

	if x == 0{
		if len(slice) == 0{
			return 0
		} else {
			for i := range slice{
				if slice[i] >= value {
					return i 
				} 
				if i == len(slice) - 1 && slice[i] < value{
					return i + 1
				}
			}
		}
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
