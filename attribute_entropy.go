/*
Marcelo
*/
package main

import (
	"fmt"
)

func attribute_entropy(ds dataset, label_collum string, attribute string) float64 {

	if ds.columns[attribute].is_numerical {
		return -1
	}

	column := ds.getColumn(attribute)
	possible_values := make(map[string]int)
	for i := 0; i < len(column); i++ {
		_, ok := possible_values[fmt.Sprintf("%v", column[i])]
		if !ok {
			possible_values[fmt.Sprintf("%v", column[i])] = 1
		} else {
			possible_values[fmt.Sprintf("%v", column[i])] += 1
		}
	}

	subsets := make([]*dataset, 0)
	entropies := make([]float64, 0)
	for key, _ := range possible_values {
		print(key)
		sub := ds.filter(func(entry *line) bool { return entry.getColumn(attribute) == key })
		subsets = append(subsets, sub)
		entropies = append(entropies, set_entropy(sub))
	}

	weighted_sum := 0.
	for i := 0; i < len(subsets); i++ {
		weighted_sum += entropies[i] * (float64(len(subsets[i].data) / len(ds.data)))
	}

	return weighted_sum
}

/*
	DEBUG - Anotações
	Valores em entropies estão como NaN
	Não entendi o funcionamento da set_entropy, a label_collum não deveria ser utilizada?
*/
