package main

import "fmt"

func main() {
	var A [10]int
	var B [10]int
	var C [10]int

	fmt.Println("Ilk dizinin 10 elemanini giriniz: ")
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &A[i])
	}

	fmt.Println("ilk dizi:", A)
	gecici := 0
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 10; j++ {
			if A[i] < A[j] {
				gecici = A[i]
				A[i] = A[j]
				A[j] = gecici
			}
		}
	}
	fmt.Println(A)

	fmt.Println("Ikinci dizinin 10 elemanini giriniz: ")
	for j := 0; j < 10; j++ {
		fmt.Scanf("%d", &B[j])
	}
	fmt.Println("ikinci dizi:", B)
	gecici = 0
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 10; j++ {
			if B[i] > B[j] {
				gecici = B[i]
				B[i] = B[j]
				B[j] = gecici
			}
		}
	}
	fmt.Println(B)

	for i := 0; i < 10; i++ {
		C[i] = A[i] + B[i]
	}
	fmt.Println("C dizisi: ", C)

}
