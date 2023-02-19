package main

import (
        "fmt"
        "github.com/faustinoaq/goworker"
)

func init() {
        goworker.Register("Resolve", resolve)
}

func main() {
        path := getResolverPath()
        fmt.Println("Resolver Path:", path)
        if err := goworker.Work(); err != nil {
                fmt.Println("Error:", err)
        }
}
