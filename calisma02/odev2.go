package main

import "fmt"

func main() {

	var sayi int

	for {
		fmt.Println("1-7 arasinda bir sayi giriniz:\nCikmak icin 0 giriniz")
		fmt.Scanf("%d", &sayi)

		if sayi == 0 {
			break
		}

		switch sayi {
		case 1:
			fmt.Println("Pazartesi")
		case 2:
			fmt.Println("Sali")
		case 3:
			fmt.Println("Carsamba")
		case 4:
			fmt.Println("Persembe")
		case 5:
			fmt.Println("Cuma")
		case 6:
			fmt.Println("Cumartesi")
		case 7:
			fmt.Println("Pazar")
		default:
			fmt.Println("--Gecersiz bir sayi girdiniz.--")
		}
	}

}
