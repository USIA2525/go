package main

import "fmt"


func main() {

     a1 := [3]int{1, 2, 3}
     a2 := [3]int{4, 5, 6}

     a1=a2

     fmt.Printf("a1 after assingment: %v\n", a1)

     a1[0] = 8
     a1[1] = 1
     a1[2] = 0

     fmt.Printf("a1 after update: %v\n", a1)
     fmt.Printf("a2 stays same: %v\n", a2)

}
