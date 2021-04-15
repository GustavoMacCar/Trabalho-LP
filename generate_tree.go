/*
Amélia
*/
package main

type split struct{
	attr string
	pos float32
	gain float32
}

type tree struct{
	split string
	children []tree
}
/*
func generate_tree(ds dataset,thresh float32, label string){
	split = best_split(ds)
	if ()
}*/