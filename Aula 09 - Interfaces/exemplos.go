package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64 // retorna um método
}

func escreverArea(f forma) {
	fmt.Println("A área da forma é %0.2f", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
	//math.Pow(c.raio, 2)
}

//Tipo genérico? xD

func generica(interf interface{}) {
	println(interf)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "String",
		float64(100): true,
		"String":     "String",
		true:         float64(12),
	}

	println(mapa)
}
