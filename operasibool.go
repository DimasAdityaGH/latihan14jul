package main

import "fmt"

func main () {
	var ujian = 100
	var magang = 90

	var hasilUjian = ujian >= 100
	var hasilMagang = magang > 80

	var hasil = hasilUjian || hasilMagang
	fmt.Println(hasil)
}