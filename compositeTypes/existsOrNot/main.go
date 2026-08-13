package main

import "fmt"

func CheckStock(inventory map[string]int, sku string) string {
	// Your code here
	count, ok := inventory[sku]
	if !ok {
		return "unknown SKU"
	} else if count == 0 {
		return "out of stock"
	}
	return fmt.Sprintf("%d in stock", count)
}

func Discontinue(inventory map[string]int, sku string) bool {
	// Your code here
	_, ok := inventory[sku]
	delete(inventory, sku)
	return ok
}

func main() {
	stock := map[string]int{
		"GO-SHIRT":   42,
		"GO-MUG":     0,
		"GO-STICKER": 100,
	}

	fmt.Println(CheckStock(stock, "GO-SHIRT"))
	fmt.Println(CheckStock(stock, "GO-MUG"))
	fmt.Println(CheckStock(stock, "GO-PLUSHIE"))

	fmt.Println(Discontinue(stock, "GO-MUG"))
	fmt.Println(Discontinue(stock, "GO-MUG"))
}
