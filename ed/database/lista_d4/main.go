
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
// 	ll := NewLList[int]()

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
			// fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "clear":
			// ll.Clear()
		case "forward":
			// search, _ := strconv.Atoi(args[1])
			// steps, _ := strconv.Atoi(args[2])
			// node := ll.Search(search)
			// if node == nil {
			// 	fmt.Println("fail: valor não encontrado")
			// 	continue
			// }
			// collect := []string{}
			// for range steps {
			// 	collect = append(collect, fmt.Sprintf("%v", node.Value))
			// 	node = node.Next()
			// }
			// fmt.Printf("[ %s ]\n", strings.Join(collect, " "))
		case "backward":
			// search, _ := strconv.Atoi(args[1])
			// steps, _ := strconv.Atoi(args[2])
			// node := ll.Search(search)
			// if node == nil {
			// 	fmt.Println("fail: valor não encontrado")
			// 	continue
			// }
			// collect := []string{}
			// for range steps {
			// 	collect = append(collect, fmt.Sprintf("%v", node.Value))
			// 	node = node.Prev()
			// }
			// fmt.Printf("[ %s ]\n", strings.Join(collect, " "))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
