/*
* @Author: GeekWho
* @Date:   2018-10-27 17:22:52
* @Last Modified by:   GeekWho
* @Last Modified time: 2018-11-18 15:52:35
*/
package main

import (
    "fmt"
    "sync"
    //"time"
)

type Fetcher interface {
    // Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
    // TODO: 并行的抓取 URL。
    // TODO: 不重复抓取页面。
    // 下面并没有实现上面两种情况：
    if depth <= 0 {
        return
    }
    if validUrl(url) {
        return
    }
    markUrl(url , 1)
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    //开启协程处理二级URL
    ch := make(chan int, len(urls))
    for _, u := range urls {
        go crawl(u, depth-1, fetcher, ch)
    }
    for i:=0; i < len(urls); i++ {
         <- ch
    }

    return
}
//内部并行，通过协程处理抓取
func crawl(url string, depth int, fetcher Fetcher, ch chan int) {
    Crawl(url, depth, fetcher)
    ch <- 1
}

//定义SafeCounter结构体，类似计数器有两个变量，一个是字符串和整数的映射，另外一个是互斥锁的实现。
type SafeCounter struct {
    v   map[string]int
    mux sync.Mutex
}

//定义全局变量
var sc SafeCounter

func validUrl(url string)bool{
    sc.mux.Lock()
    _ , ok := sc.v[url]
    defer sc.mux.Unlock()
    return ok
}

func markUrl(url string, count int) {
    sc.mux.Lock()
    defer sc.mux.Unlock()
    sc.v[url] += count
}

func main() {
    sc = SafeCounter{v: make(map[string]int)}
    Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
    "https://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "https://golang.org/pkg/",
            "https://golang.org/cmd/",
        },
    },
    "https://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "https://golang.org/",
            "https://golang.org/cmd/",
            "https://golang.org/pkg/fmt/",
            "https://golang.org/pkg/os/",
        },
    },
    "https://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
    "https://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
}

