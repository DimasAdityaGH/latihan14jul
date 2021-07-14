package main

import "fmt"

func main () {
	var produk [3]string
	produk[0] = "susu"
	produk[1] = "gula"
	produk[2] = "air"
	fmt.Println(produk)

	var harga = [03]int16{
		10000,
		20000,
		10000,
	}
	fmt.Println(harga[0])
}