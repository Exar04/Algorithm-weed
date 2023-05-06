package main

import "fmt"

func multiply(n int) {
    for addThisManyTimes := n; addThisManyTimes > 0; addThisManyTimes-- {
        n = n + n
    }

} 

func main() {
    count := 1
    for i:= 10000000000; i < 100000000000000000 ; i++{
        multiply(i)
        fmt.Println(count)
        count++
    }
}
