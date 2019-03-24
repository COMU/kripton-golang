package main

import "fmt"

func main() {
	sonuc := 1
	var dizi [20]int
	dizi[0] = 1
	dizi[1] = 2
	for i := 2; i <= 19; i++ {
		sonuc = sonuc + dizi[i-1]
		dizi[i] = sonuc + 1

	}

	fmt.Println(dizi)

}
