/*
* @Author: GeekWho
* @Date:   2018-10-27 17:18:22
* @Last Modified by:   GeekWho
* @Last Modified time: 2018-10-27 18:19:22
*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func fibonacci_1(x int) int {
    if x == 1 {
        return 1
    }
    if x > 1 {
        return fibonacci_1(x-2) + fibonacci_1(x-1)
    }
    return 0
}

func main() {
    f := fibonacci()
    for i := 1; i <= 10; i++ {
        fmt.Println(f())
    }

    //计算第几个数字
    for i := 0; i < 10; i++ {
        fmt.Println(fibonacci_1(i))
    }
}
