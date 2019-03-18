package main

import "fmt"

func main() {
	var sayi int = 1
	for sayi != 0 {
		fmt.Println("Bir sayi giriniz >>>")
		fmt.Scanf("%d", &sayi)
		switch sayi {
		case 0:
			fmt.Println("güle güle!!!")
		case 1:
			fmt.Println("1. gün pazartesi")
		case 2:
			fmt.Println("2. gün sali")
		case 3:
			fmt.Println("3. gün carsamba")
		case 4:
			fmt.Println("4. gün persembe")
		case 5:
			fmt.Println("5. gün cuma")
		case 6:
			fmt.Println("6. gün cumartesi")
		case 7:
			fmt.Println("7. gün pazar")
		default:
			fmt.Println("1-7 arasında bir deger girmelisiniz!")
		}
	}

}
