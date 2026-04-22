package main

import (
  "fmt"
  "math"


)


func doSomething(a, b uint32) bool {
    if (math.MaxUint32 - a) < b {
         fmt.Println("error")
         return false
    } else {

          sum := a + b

          fmt.Println("n =", sum)


          return true
    }
}


func main() {
         fmt.Print("test1: ")
         doSomething(100, 100)

         fmt.Print("test2: ")
         doSomething(math.MaxUint32-10, 5)

         fmt.Print("test3: ")
         doSomething(math.MaxUint32-10, 15)

}
