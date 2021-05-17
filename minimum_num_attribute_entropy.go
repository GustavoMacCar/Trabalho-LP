/*
Gustavo
*/
package main

import "errors"

//"fmt"

func minimum_num_attribute_entropy(ds *dataset, label_column string, attribute string) (float64, float32, error) {

	if !ds.columns[attribute].is_numerical {
		return -1, -1, errors.New("attribute is not numerical")
	}

	attribute_values := ds.getColumn(attribute)

	min_attribute_value := attribute_values[0]
	max_attribute_value := attribute_values[0]
	for _, v := range attribute_values {
		if v.(float32) < min_attribute_value.(float32) {
			min_attribute_value = v
		}

		if v.(float32) > max_attribute_value.(float32) {
			max_attribute_value = v
		}
	}

	initial_threshold := (min_attribute_value.(float32) + max_attribute_value.(float32)) / 2

	initial_entropy, _ := num_attribute_entropy(ds, label_column, attribute, initial_threshold)

	entropy_gt := initial_entropy
	entropy_gt_new := initial_entropy
	threshold_gt := initial_threshold

	entropy_lt := initial_entropy
	entropy_lt_new := initial_entropy
	threshold_lt := initial_threshold

	for {
		threshold_gt += 0.1
		entropy_gt_new, _ = num_attribute_entropy(ds, label_column, attribute, threshold_gt)
		if entropy_gt_new >= entropy_gt {
			break
		}
		entropy_gt = entropy_gt_new
	}

	for {
		threshold_lt -= 0.1
		entropy_lt_new, _ = num_attribute_entropy(ds, label_column, attribute, threshold_lt)
		if entropy_lt_new >= entropy_lt {
			break
		}
		entropy_lt = entropy_lt_new
	}

	if entropy_lt < entropy_gt {
		return entropy_lt, threshold_lt, nil
	}

	return entropy_gt, threshold_gt, nil
	//return 0, 0, nil

}
