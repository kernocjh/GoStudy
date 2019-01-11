package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    resp1, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    nbytes1, err := io.Copy(ioutil.Discard, resp1.Body)
    resp.Body.Close()
    resp1.Body.Close()


    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    if nbytes != nbytes1 {
        ch <- fmt.Sprintf("getting urls is not same by tiwces\n")
        return
    }
    
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
