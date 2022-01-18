package main

import "fmt"

func main() {
	var buah = []string{"Apel", "Jeruk", "Anggur", "Melon", "Pepaya"}
	var j = len(buah)
	fmt.Println("Jumlah Element : ", j)
	fmt.Println("Isi Elemen : ", buah)
	var i = 0
	for i < j {
		fmt.Println("Element Ke - : ", i, buah[i])
		i++
	}
}
