package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	tbl := NewTable(file)
	tbl.Optimize()

	fmt.Println("Table", tbl, "has total happiness", tbl.TotalHappiness())

	tbl = append(tbl, Guest{Name: "Host"})
	tbl.Optimize()

	fmt.Println("Table", tbl, "has total happiness", tbl.TotalHappiness())
}
