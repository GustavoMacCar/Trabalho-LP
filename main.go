package main

import (
	"fmt"
)

func testDs() {
	fmt.Println("Loading data...")
	sq := loadDataset("squares.csv")
	//fmt.Println(sq)
	fmt.Println("First 5 lines of the dataset:")
	sq.print(5)
	fmt.Println("\n")

	fmt.Println("First 5 lines' \"x\" column")
	fmt.Println(sq.getColumn("colour")[:5], "\n")

	fmt.Println("First line:\n")
	sq.data[0].print()
	fmt.Println("\n")

	fmt.Println("Column \"a\" of the first line (doesn't exist)")
	fmt.Println("\t", sq.data[0].getColumn("a"))
	fmt.Println("\nColumn \"x\" of the first line")
	fmt.Println("\t", sq.data[0].getColumn("x"))
	fmt.Println("\n")

	orange := sq.filter(func(entry *line) bool { return entry.getColumn("colour") == "orange" })
	fmt.Println("First 5 lines of the sub-dataset containing only the orange entries")
	orange.print(5)
	fmt.Println("\n")

	fmt.Println("Total orange:", len(orange.data))
	fmt.Println("Total blue:", sq.count(func(entry *line) bool { return entry.getColumn("colour") == "blue" }))

	res := make(chan *tree)
	go generate_tree(sq, "colour", 0.4, 2, 0, res)
	classifier := <-res
	print_tree(classifier, 0)
}

func main() {
	testDs()
}
