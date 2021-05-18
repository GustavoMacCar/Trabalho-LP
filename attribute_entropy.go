/*
Marcelo
*/
package main

import "errors"

func attribute_entropy(ds *dataset, label_column string, attribute string) (float64, error) {
	/*

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
			entropies = append(entropies, set_entropy(sub, attribute))
		}

		weighted_sum := 0.
		for i := 0; i < len(subsets); i++ {
			weighted_sum += entropies[i] * (float64(len(subsets[i].data) / len(ds.data)))
		}

		return weighted_sum
		//return 0 */

	if ds.columns[attribute].is_numerical {
		return -1, errors.New("attribute is not categorical")
	}

	var total int = len(ds.data)

	var possible_values []string
	var groups []*dataset
	var entropy float64 = 0

	for _, e := range ds.data {
		if !contains(possible_values, e.getColumn(attribute).(string)) {
			possible_values = append(possible_values, e.getColumn(attribute).(string))
		}
	}

	for _, e := range possible_values {
		groups = append(groups, ds.filter(func(entry *line) bool { return entry.getColumn(attribute) == e }))
	}

	for i, _ := range groups {
		current_length := groups[i].count(func(entry *line) bool { return true })
		entropy += set_entropy(groups[i], label_column) * (float64(current_length) / float64(total))
	}

	return entropy, nil

}

/*
	DEBUG - Anotações
	Valores em entropies estão como NaN
	Não entendi o funcionamento da set_entropy, a label_collum não deveria ser utilizada?
*/
