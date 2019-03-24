//her elemani, kendinden onceki elemanlarin toplaminin 1 fazlasi olarak tanimlanmis dizi
package main

import "fmt"

func main() {
	array := [20]int{
		1,
		2,
	}

	for i := 2; i < 20; i++ {
		for j := 0; j < i; j++ {
			array[i] += array[j]
		}
		array[i] += 1

	}
	fmt.Println(array)

}
