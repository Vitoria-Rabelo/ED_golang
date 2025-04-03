package main
import "fmt"
import "math"

func calcularArea(num1, num2, num3 float64) float64{
    var perimetro float64 = (num1 + num2 + num3) / 2
    var area float64 = math.Sqrt(perimetro * (perimetro - num1) * (perimetro - num2) * (perimetro - num3))
    return area
}

func main() {
    var a, b, c float64

    fmt.Scan(&a, &b, &c)

    var area float64 = calcularArea(a, b, c)
    fmt.Printf("%.2f\n", area)
}
