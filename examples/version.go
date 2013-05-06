package main

import "fmt"
import curl "github.com/edmonds/golang-curl"

func main() {
    fmt.Println(curl.Version())
    curl.GlobalCleanup()
}
