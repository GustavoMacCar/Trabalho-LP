/*
Daniel
*/
package main

import (
	"math"
)

func set_entropy(ds *dataset, attribute string) float64 {
	/*
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

		return entropy */

	var categories []string
	var entropy float64
	var total float64

	entropy = 0

	for _, e := range ds.data {
		if !contains(categories, e.getColumn(attribute).(string)) {
			categories = append(categories, e.getColumn(attribute).(string))
		}
		total += 1
	}

	for _, e := range categories {
		current_category_count := float64(ds.count(func(entry *line) bool { return entry.getColumn(attribute) == e }))
		entropy += -math.Log2(current_category_count/total) * current_category_count / total

	}

	//fmt.Println(entropy)

	return entropy
}
