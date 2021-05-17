/*
Gustavo
num_attribute_entropy(data:pd.DataFrame,label_column:str,attribute:str,threshold:float)->float
*/
package main

import (
	"errors"
)

func num_attribute_entropy(ds *dataset, label_column string, attribute string, threshhold float32) (float64, error) {
	if !ds.columns[attribute].is_numerical {
		return -1, errors.New("attribute is not numerical")
	}

	below := ds.filter(func(entry *line) bool { return entry.getColumn(attribute).(float32) < threshhold })

	above := ds.filter(func(entry *line) bool { return entry.getColumn(attribute).(float32) >= threshhold })

	//entropy_below := set_entropy(below.getColumn(label_column))
	entropy_below := set_entropy(below, label_column)
	//entropy_above := set_entropy(above.getColumn(label_column))
	entropy_above := set_entropy(above, label_column)

	total_length := float64(len(ds.data))

	below_length := float64(len(below.data))

	above_length := float64(len(above.data))

	att_entropy := (entropy_below*(below_length/total_length) + entropy_above*(above_length/total_length))
	//fmt.Println(att_entropy)
	return att_entropy, nil
	//return 0.0, nil

}
