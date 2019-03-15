package main

import "fmt"

func main() {
    var sayi int
    fmt.Println("1 ile 7 arasında bir sayı giriniz")
    fmt.Scanf("%d", &sayi)

    for sayi != 0 {
        switch sayi {
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
            fmt.Println("Geçersiz bir sayı girdiniz")

        }
        fmt.Println("1 ile 7 arasında bir sayı giriniz")
        fmt.Scanf("%d", &sayi)

    }

}

