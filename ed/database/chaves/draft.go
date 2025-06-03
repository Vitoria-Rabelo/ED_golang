package main
import "fmt"

 type Node struct{
    time string
    next *Node
}

type Queue struct{
        head *Node
        tail *Node
    }

func (q *Queue) Enqueue(time string) {
    newNode := &Node{
        time: time,
        next: nil,
    }

    if q.tail == nil {
        q.head = newNode
        q.tail = newNode
    }
    q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) Dequeue() string{
    if q.head == nil{
        return ""
    }

    team := q.head.time
    q.head = q.head.next

    if q.head == nil{
        q.tail = nil
    }
    return team
}

func main() {
    queue := Queue{}

	for _, time := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"} {
		queue.Enqueue(time)
	}

    for i := 0; i < 15; i++ {
		var m, n int
		fmt.Scan(&m, &n)

        t1 := queue.Dequeue()
		t2 := queue.Dequeue()

        if m > n {
			queue.Enqueue(t1)
		} else {
			queue.Enqueue(t2)
		}
    }

    fmt.Println(queue.Dequeue())

}