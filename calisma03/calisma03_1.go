package main

import "fmt"

func main() {
	var A [10]int
        var B [10]int
        var C [10]int
        fmt.Println("A dizisi için 10 adet sayı girin.\n")
        for i := 0; i < 10; i++ {
                fmt.Scanf("%d", &A[i])
        }
        fmt.Println("B dizisi için  10 adet sayı girin.\n")
        for i := 0; i < 10; i++ {
                fmt.Scanf("%d", &B[i])
        }
        fmt.Println("\n")
        for i := 0; i < 10; i++ {
                C[i] = A[i] + B[i]
                fmt.Println(C[i])
        }
}
