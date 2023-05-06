package main

import "fmt"

func multiply(n int) {
    for addThisManyTimes := n; addThisManyTimes > 0; addThisManyTimes-- {
        n = n + n
    }

} 

func main() {
    slowItMore := 0 
    fmt.Print("Enter a number to make it more slow - ")
    fmt.Scanf("%d", &slowItMore)
    fmt.Println()
    count := 1
    for i:= 10000000000 ; i < 100000000000000000 ; i++{
        multiply(i* slowItMore)
        fmt.Println(count)
        count++
    }
}
