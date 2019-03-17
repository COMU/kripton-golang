package main
  
import "fmt"

func main() {
        var sayi int
        fmt.Println("1 ile 7 arasında bir sayı giriniz.Çıkmak için 0 giriniz.\n")
        fmt.Scanf("%d", &sayi)

        for sayi != 0 {
                switch sayi {
                case 1:
                        fmt.Println("Pazartesi\n")
                        fmt.Println("1 ile 7 arasında bir sayı giriniz.\n")
                        fmt.Scanf("%d", &sayi)
                case 2:
                        fmt.Println("Salı\n")
                        fmt.Println("1 ile 7 arasında bir sayı giriniz.\n")
                        fmt.Scanf("%d", &sayi)
                case 3:
                        fmt.Println("Çarşamba\n")
                        fmt.Println("1 ile 7 arasında bir sayı giriniz.\n")
                        fmt.Scanf("%d", &sayi)
                case 4:
                        fmt.Println("Perşembe\n")
                        fmt.Println("1 ile 7 arasında bir sayı giriniz.\n")
                        fmt.Scanf("%d", &sayi)
                case 5:
                        fmt.Println("Cuma\n")
                        fmt.Println("1 ile 7 arasında bir sayı giriniz.\n")
                        fmt.Scanf("%d", &sayi)
                case 6:
                        fmt.Println("Cumartesi\n")
                        fmt.Println("1 ile 7 arasında bir sayı giriniz.\n")
                        fmt.Scanf("%d", &sayi)
                case 7:
                        fmt.Println("Pazar\n")
                        fmt.Println("1 ile 7 arasında bir sayiriniz.\n")
                        fmt.Scanf("%d", &sayi)
                default:
                        fmt.Println("\t Geçersiz bir sayı girdiniz. \t \n \t Lütfen tekrar bir sayı giriniz.\t")
                        fmt.Scanf("%d", &sayi)
                }
        }
}







