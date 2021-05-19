package main

import (
	"math"
)

func bestSplit(ds *dataset, label_column string) (string, float64, float64) {
	/*

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
		return "", 0, 0 */

	//	initial_entropy := set_entropy(ds, label_column)
	current_entropy := set_entropy(ds, label_column)
	var attributes []string

	for column := range ds.columns {
		if column != label_column {
			attributes = append(attributes, column)
		}
	}

	var entropies []float64
	var thresholds []float64
	for _, attribute := range attributes {
		if !ds.columns[attribute].is_numerical {
			current_entropy, _ := attribute_entropy(ds, label_column, attribute)
			entropies = append(entropies, current_entropy)
			thresholds = append(thresholds, float64(math.Inf(-1)))
		} else {
			entropy, threshold, _ := minimum_num_attribute_entropy(ds, label_column, attribute)
			entropies = append(entropies, entropy)
			thresholds = append(thresholds, threshold)
		}
	}

	var gains []float64
	for _, ent := range entropies {
		gains = append(gains, (current_entropy - ent))
	}
	var biggest float64 = gains[0]
	var index int = 0

	for i, _ := range gains {
		if gains[i] > biggest {
			biggest = gains[i]
			index = i
		}
	}

	var final_attribute string
	var final_threshold float64

	final_attribute = attributes[index]
	final_threshold = thresholds[index]

	return final_attribute, float64(final_threshold), biggest

}
