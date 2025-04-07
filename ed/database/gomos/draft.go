package main
import "fmt"

func main() {
    var q int
    var d string
    fmt.Scan(&q, &d)

    var x [100]int
    var y [100]int

    for i := 0; i < q; i++ {
        fmt.Scan(&x[i], &y[i])
    }

    oldX, oldY := x[0], y[0]

    switch d {
    case "L":
        x[0]--
    case "R":
        x[0]++
    case "U":
        y[0]--
    case "D":
        y[0]++
    }

    for i := 1; i < q; i++ {
        tempX, tempY := x[i], y[i]
        x[i], y[i] = oldX, oldY
        oldX, oldY = tempX, tempY
    }

    for i := 0; i < q; i++ {
        fmt.Println(x[i], y[i])
    }
}
