package main

import (
    "io"
    "strconv"
    "fmt"
    "net/http"
    "os"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    baseURL := "https://yys.res.netease.com/pc/zt/20161108171335/data/shishen_big_beforeAwake/"

    for i := 200; i <= 401; i++ {
        wg.Add(1)
        go func(i int) {
            image := strconv.Itoa(i) + ".png"
            imageURL := baseURL + image
            
            fmt.Println("Starting to download for " + image)

            response, err := http.Get(imageURL)
            checkErr(err)
            defer response.Body.Close()

            if response.StatusCode >= 200 && response.StatusCode <= 299 {
                file, err := os.Create("../resources/original/" + image)
                checkErr(err)
                
                _, err = io.Copy(file, response.Body)
                checkErr(err)
                file.Close()
            }

            wg.Done()
            fmt.Println("Finished downloading for " + image)
        }(i)
    }

    wg.Wait()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}