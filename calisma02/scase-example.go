package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("1 ile 7 arasında bir rakam giriniz:")
	fmt.Scanf("%d", &number)

	for number != 0 {
		switch number {
		case 1:
			fmt.Println("Pazartesi")
		case 2:
			fmt.Println("Salı")
		case 3:
			fmt.Println("Çarşamba")
		case 4:
			fmt.Println("Perşembe")
		case 5:
			fmt.Println("Cuma")
		case 6:
			fmt.Println("Cumartesi")
		case 7:
			fmt.Println("Pazar")
		default:
			fmt.Println("Geçerli rakam girin.")
		}
		fmt.Println("1 ile 7 arasında bir rakam giriniz.")
		fmt.Scanf("%d", &number)

	}
}
