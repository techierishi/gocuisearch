package main

import (
	"fmt"
	"gocuisearch/gocuisearch"
	"os"
	"strconv"
)

const exitCodeExecute = 111

func main() {

	var records []gocuisearch.SearchApp
	for i := 0; i < 2000; i++ {
		records = append(records, gocuisearch.NewSearchApp(&gocuisearch.RowItem{
			Content: `
				kubectl get pod` + strconv.Itoa(i),
		}))
	}

	output, idx, _ := gocuisearch.CuiSearch(records)
	fmt.Print(output, idx)
	os.Exit(0)
}
