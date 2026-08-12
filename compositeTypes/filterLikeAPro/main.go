package main

import (
	"fmt"
	"slices"
)

func DropTempTables(tables []string) []string {
	// Your code here
	tables = slices.DeleteFunc(tables, func(table string) bool {
		if len(table) < 6 {
			return false
		}
		matchStr := table[:5]
		return matchStr == "temp_"
	})
	return tables
}

func main() {
	tables := []string{"users", "temp_001", "orders", "temp_cache", "products"}
	cleaned := DropTempTables(tables)
	fmt.Println(cleaned)
}
