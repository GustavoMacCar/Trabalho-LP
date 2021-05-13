package main

import "sort"

func bestSplit(ds *dataset, label_column string) (string, float64, float64) {
	entropy := set_entropy(ds)

	//attributes = [column for column in data.columns if column!=label_column]
	attributes := make([]string, 0)

	for column := range ds.columns {
		if column != label_column {
			attributes = append(attributes, column)
		}
	}

	// entropies = [(attribute_entropy(data,label_column,attribute),None) if data[attribute].dtype == object else minimum_num_attribute_entropy(data,label_column,attribute) for attribute in attributes]
	entropies := make([]float64, 0)
	thresholds := make([]float64, 0)
	for attribute := range attributes {
		if !ds.columns[label_column].is_numerical {
			entropies = append(entropies, attribute_entropy(ds, label_column, attribute))
			thresholds = append(thresholds, 0)
		} else {
			entropy, threshold, _ := minimum_num_attribute_entropy(ds, label_column, attribute)
			entropies = append(entropies, entropy)
			thresholds = append(thresholds, threshold)
		}
	}

	//gains = [(entropy-ent[0],i) for i,ent in enumerate(entropies)]
	gains := make([]float64, 0)
	for i, ent := range entropies {
		gains = append(gains, (entropy - ent))
		index := i
	}

	//minimum = sorted(gains,reverse=True)[0]
	sort.Sort(sort.Reverse(sort.Float64Slice(gains)))
	minimum := gains[0]

	//return (attributes[minimum[1]],entropies[minimum[1]][1],minimum[0])
	return attributes[int(minimum)], entropies[int(minimum)], minimum
}
