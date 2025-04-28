package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func tree(p *Pen, tam float64) {

	if tam < 5 {
		if randInt(0, 100) < 5{
			p.SetColor(255,0,0)
			p.DrawCircle(5)
			p.SetColor(0,0,0)
		}
		return
	}

	p.SetStrokeWidth(int(tam) / 5)
	p.Walk(tam)
	p.Rotate(35.0)
	tree(p, tam*0.75)
	p.Rotate(-2*35)
	tree(p, tam*0.75)
	p.Rotate(35.0)
	p.Walk(-tam)
}

func embua (p *Pen, tam float64){
	if tam < 10{
		return
	}

	p.SetColor(randInt(0, 255),randInt(0, 255),randInt(0, 255) )
	p.Walk(tam)
	p.Rotate(-90)
	embua(p, tam - 5)
}

func circle (p *Pen, tam float64){
	if tam < 1{
		return
	}
	p.SetDown(true)
	p.DrawCircle(5)
	p.SetDown(false)
	for range 6 {
		p.Walk(tam)
		circle(p, tam / 3)
		p.Walk(-tam)
		p.Rotate(60)
	}

}

func frozen (p *Pen, tam float64){
	if tam < 1{
		return
	}

	for range 5 {
		p.Walk(tam)
		frozen(p, tam / 3)
		p.Walk(-tam)
		p.Rotate(360/5)
	}

}

func main() {
	pen := NewPen(500, 500)
	pen.SetPosition(250, 250)
	pen.SetHeading(90)
	//embua(pen, 460)
	//tree(pen, 100)
	//circle(pen, 180)
	frozen(pen, 180)
	pen.SaveToFile("tree.svg")
	fmt.Println("SVG file created successfully.")
}
