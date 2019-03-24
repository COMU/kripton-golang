package main

import "fmt"

func main() {
        var list [20]int
        list[0] = 1
        list[1] = 2
        carpim := 2
        for i := 2; i < 20; i++ {
                list[i] = carpim + 1
                carpim *= list[i]
        }
        for i := 0; i < 20; i++ {
                fmt.Println(list[i])
        }
}
