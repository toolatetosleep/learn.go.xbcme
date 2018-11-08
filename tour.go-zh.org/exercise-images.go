/*
* @Author: GeekWho
* @Date:   2018-10-27 17:21:52
* @Last Modified by:   GeekWho
* @Last Modified time: 2018-11-08 23:56:01
*/
package main

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
