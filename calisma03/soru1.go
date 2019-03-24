package main

import "fmt"

func main() {
	var A [10]int
	var B [10]int
	var C [10]int
	for i := 0; i <= 9; i++ {
		fmt.Println("Bir sayı giriniz")
		fmt.Scanf("%d", &A[i])

	}
	for m := 0; m < len(A); m++ {
		for n := 0; n < len(A)-1-m; n++ {
			if A[n] > A[n+1] {
				A[n], A[n+1] = A[n+1], A[n]
			}
		}
	}

	fmt.Println(A)

	for j := 0; j <= 9; j++ {
		fmt.Println("Bir sayı giriniz")
		fmt.Scanf("%d", &B[j])

	}

	for k := 0; k < len(B); k++ {
		for l := 0; l < len(B)-1-k; l++ {
			if B[l] < B[l+1] {
				B[l], B[l+1] = B[l+1], B[l]
			}
		}
	}

	fmt.Println(B)
	for p := 0; p <= 9; p++ {

		C[p] = A[p] + B[p]

	}
	fmt.Println(C)

}
