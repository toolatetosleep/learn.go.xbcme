/*
* @Author: GeekWho
* @Date:   2018-10-27 17:20:59
* @Last Modified by:   GeekWho
* @Last Modified time: 2018-10-28 17:36:06
*/
package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return -2,ErrNegativeSqrt(-2)
    }
    z := 1.0
     for i :=1;(z - x) < 0.001 && i < 10;i++ {
        z -= (z*z - x) / (2*z)
     }
     return z, nil
}
func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
