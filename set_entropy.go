/*
Daniel
*/
package main

import (
	"fmt"
	"math"
)

type classe struct {
	tag      string
	contagem float64
}

func set_entropy(ds *dataset) float64 {
	var azul float64 = 0
	var laranja float64 = 0
	var total float64 = 0
	//file := loadDataset("squares.csv")
	azul = float64(ds.count(func(entry *line) bool { return entry.getColumn("colour") == "blue" }))
	laranja = float64(ds.count(func(entry *line) bool { return entry.getColumn("colour") == "orange" }))
	fmt.Print(azul, "\n")
	fmt.Print(laranja, "\n")
	total = azul + laranja
	var entropy = -math.Log2(azul/total)*azul/total - math.Log2(laranja/total)*laranja/total
	//fmt.Print(entropy)
	//definição dinâmica de classes existentes
	//slice := make([]classe, 1)

	return entropy
}
