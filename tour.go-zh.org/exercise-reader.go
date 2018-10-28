/*
* @Author: GeekWho
* @Date:   2018-10-27 17:21:16
* @Last Modified by:   GeekWho
* @Last Modified time: 2018-10-28 19:07:14
*/
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
    for i := range b {
        b[i] = 'A'
    }
    return len(b), nil
}

func main() {
    reader.Validate(MyReader{})
}
