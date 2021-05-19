/*
Amélia
*/
package main

import (
	"fmt"
)

type split struct {
	attr   string
	pos    float64
	gain   float64
	fields []string
}

type tree struct {
	split    split
	children []*tree
}

func getAttr(ds *dataset, field string) map[string]bool {
	attributes := map[string]bool{}
	for _, i := range ds.getColumn(field) {
		attributes[i.(string)] = true
	}
	return attributes
}

func leaf(ds *dataset, label_column string) string {
	attributes := getAttr(ds, label_column)
	var maxname string
	maxcount := 0
	for attr, _ := range attributes {
		count := ds.count(func(entry *line) bool { return entry.getColumn(label_column) == attr })
		if count > maxcount {
			maxname = attr
			maxcount = count
		}
	}
	return maxname
}

func print_tree(t *tree, lv int) {
	for i := 0; i < lv; i++ {
		fmt.Print("\t")
	}
	if len(t.children) == 0 {
		fmt.Printf("It's %s\n", t.split.attr)
	} else {
		fmt.Printf("Split on %s at %.2f\n", t.split.attr, t.split.pos)
		for i := 0; i < lv; i++ {
			fmt.Print("\t")
		}
		fmt.Print("If greater than\n")
		print_tree(t.children[0], lv+1)
		for i := 0; i < lv; i++ {
			fmt.Print("\t")
		}
		fmt.Print("Otherwise\n")
		print_tree(t.children[1], lv+1)
	}
}

func generate_tree(ds *dataset, label_column string, gain_thr float64, min_len int, cur_len int, c chan *tree) {
	attr, pos, gain := bestSplit(ds, label_column)
	t := new(tree)

	if len(ds.data) <= 2 || (gain < gain_thr && cur_len > min_len) {
		l := leaf(ds, label_column)
		t.split.attr = l
		c <- t
		return
	}

	t.split.attr = attr
	t.split.pos = pos
	t.split.gain = gain
	fmt.Printf("Attr is %s Pos %.2f Gain %.2f\n", attr, float64(pos), gain)
	if !ds.columns[attr].is_numerical {
		//Categórico
		vals := getAttr(ds, attr)
		channels := make(map[string]chan *tree)
		for val, _ := range vals {
			channels[val] = make(chan *tree)
			go generate_tree(
				ds.filter(func(entry *line) bool {
					return entry.getColumn(attr) == val
				}),
				label_column,
				gain_thr, min_len, cur_len+1, channels[val])
		}
		for key, _ := range vals {
			t.split.fields = append(t.split.fields, key)
			tp := <-channels[key]
			t.children = append(t.children, tp)
		}
	} else {
		lt := make(chan *tree)
		gt := make(chan *tree)
		go generate_tree(
			ds.filter(func(entry *line) bool {
				return entry.getColumn(attr).(float64) > pos
			}),
			label_column,
			gain_thr, min_len, cur_len+1, lt)
		go generate_tree(
			ds.filter(func(entry *line) bool {
				return entry.getColumn(attr).(float64) <= pos
			}),
			label_column,
			gain_thr, min_len, cur_len+1, gt)
		t.children = append(t.children, <-lt, <-gt)
	}
	c <- t
}
