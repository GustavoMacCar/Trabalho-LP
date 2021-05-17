/*
Golang Dataset Utilities
Amélia O. F. da S. - 190037971
*/
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
)

/*
Dataset types
*/
type line struct {
	num       []float32
	str       []string
	cols      *map[string]column_indexer
	col_order *[]string
}

/*
Gets the value of a certain property
*/

func (l *line) getColumn(column string) interface{} {
	indexer, exists := (*l.cols)[column]
	if !exists {
		return nil
	} else if indexer.is_numerical {
		return l.num[indexer.index]
	} else {
		return l.str[indexer.index]
	}
}

/*
Prints the line
*/

func (l *line) print() {
	for _, col := range *l.col_order {
		indexer := (*l.cols)[col]
		if indexer.is_numerical {
			fmt.Print(l.num[indexer.index], " ")
		} else {
			fmt.Print(l.str[indexer.index], " ")
		}
	}
	fmt.Print("\n")
}

type column_indexer struct {
	index        int
	is_numerical bool
}

type dataset struct {
	columns   map[string]column_indexer
	col_order []string
	data      []*line
}

/*
Prints a dataset
*/

func (ds *dataset) print(i ...int) {
	var max int
	if len(i) > 0 {
		max = i[0]
	} else {
		max = len(ds.data)
	}
	for _, cname := range ds.col_order {
		fmt.Print(cname, " ")
	}
	fmt.Print("\n")
	for i, line := range ds.data {
		if i > max {
			break
		}
		line.print()
	}
}

/*
A function that can be passed to
filtering/counting functions
*/
type filterFunction func(entry *line) bool

/*
Returns a new dataset containing all lines that passed the filter function test
*/
func (ds *dataset) filter(ff filterFunction) *dataset {
	rslice := make([]*line, 0)
	for _, entry := range ds.data {
		if ff(entry) {
			rslice = append(rslice, entry)
		}
	}
	ret := dataset{
		columns:   ds.columns,
		col_order: ds.col_order,
		data:      rslice}
	return &ret
}

/*
Returns a slice containing a column of a dataset
*/

func (ds *dataset) getColumn(col string) []interface{} {
	_, exists := ds.columns[col]
	if !exists {
		return nil
	}
	rslice := make([]interface{}, len(ds.data))
	for i, line := range ds.data {
		rslice[i] = line.getColumn(col)
	}
	return rslice
}

/*
Counts the occurences of lines that pass the filter function
*/
func (ds *dataset) count(ff filterFunction) int {
	i := 0
	for _, entry := range ds.data {
		if ff(entry) {
			i++
		}
	}
	return i
}

/*
Loads a dataset from a .csv file
*/
func loadDataset(filename string) *dataset {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(bytes.NewReader(dat))
	lines, err := r.ReadAll()

	cols := make(map[string]column_indexer)
	col_order := make([]string, len(lines[0]))
	ds_lines := make([]*line, len(lines)-1)

	/*Here we try to discover which lines are numerical*/
	num_index := 0
	cat_index := 0
	for i, v := range lines[1] {
		col_order[i] = lines[0][i]
		_, err := strconv.ParseFloat(v, 32)
		if err != nil {
			cols[lines[0][i]] = column_indexer{
				index:        cat_index,
				is_numerical: false}
			cat_index++
		} else {
			cols[lines[0][i]] = column_indexer{
				index:        num_index,
				is_numerical: true}
			num_index++
		}
	}

	/*Loading lines*/
	for i, l := range lines[1:] {
		nl := line{
			num:       make([]float32, num_index),
			str:       make([]string, cat_index),
			cols:      &cols,
			col_order: &col_order}
		ds_lines[i] = &nl
		j := 0
		for column := range lines[0] {
			indexer := cols[lines[0][column]]
			if indexer.is_numerical {
				f64, _ := strconv.ParseFloat(l[j], 32)
				ds_lines[i].num[indexer.index] = float32(f64)
			} else {
				ds_lines[i].str[indexer.index] = l[j]
			}
			j++
		}
	}

	ret := dataset{
		columns:   cols,
		col_order: col_order,
		data:      ds_lines}
	return &ret
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}
