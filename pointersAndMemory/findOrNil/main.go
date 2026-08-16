package main

import "fmt"

type Item struct {
	Name  string
	Stock int
}

// LowestStock returns a pointer to the item with the least stock, so the
// caller can restock it by adding to Stock. The returned pointer must point
// into the slice, so the caller's write reaches the real item.
func LowestStock(items []Item) *Item {
	// Your code here
	minStock := items[0].Stock
	minStockItem := &items[0]
	for i := range items {
		if items[i].Stock < minStock {
			minStock = items[i].Stock
			minStockItem = &items[i]
		}
	}
	return minStockItem
}

func main() {
	items := []Item{
		{Name: "USB cable", Stock: 40},
		{Name: "HDMI cable", Stock: 6},
		{Name: "Power adapter", Stock: 22},
	}

	if it := LowestStock(items); it != nil {
		it.Stock += 50
	}

	fmt.Println(items[1].Stock) // want 56, prints 6 until LowestStock is fixed
}
