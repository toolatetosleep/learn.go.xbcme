/*
* @Author: GeekWho
* @Date:   2018-10-27 17:21:34
* @Last Modified by:   GeekWho
* @Last Modified time: 2018-11-08 01:36:45
*/
package main

import (
    "io"
    "os"
    "strings"
    "fmt"
)

type rot13Reader struct {
    r io.Reader
}
//自定义错误类型
type ErrNegativeByte string
//实现Error接口
func (e ErrNegativeByte) Error() string {
    return fmt.Sprintf("get Byte error: %v",e)
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
    //判断b的长度类型是否正确
    if len(b) <= 0 {
        return len(b),ErrNegativeByte("length is empty")
    }

    n,err = r.r.Read(b)

    /*if e != nil {
        return len(b),ErrNegativeByte(e.Error())
    }*/
    //fmt.Printf("b[:n] = %q\n", b[:n])
    for i := 0; i < n; i++ {
        b[i] = rot13(b[i])
        fmt.Printf("n =%v rot = %q Type: %T \n", n, rot13(b[i]), b[i])
    }
    return
    /*if false {
        //元素类型是否为字符串
        return len(b),ErrNegativeByte("element is no string")
    }
    return len(b), nil*/
}

func rot13(b byte) byte {
    var a, z byte
    switch {
    case 'a' <= b && b <= 'z':
        a, z = 'a', 'z'
    case 'A' <= b && b <= 'Z':
        a, z = 'A', 'Z'
    default:
        return b
    }
    return (b-a+13)%(z-a+1) + a
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
