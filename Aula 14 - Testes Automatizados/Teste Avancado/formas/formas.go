package formas

import (
	"fmt"
	"math"
)

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

type Forma interface {
	area() float64 // retorna um método
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) area() float64 {
	return math.Pi * (c.Raio * c.Raio)
	//math.Pow(c.raio, 2)
}
