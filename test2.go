package main

import (
    "fmt"
    "strings"
)

func main() {

    _srcIp := "127.0.0.2-9"
    //var srcIp []string

    srcIp := strings.Split(_srcIp, ".")
    fmt.Println(strings.Split(_srcIp, ".")[0])
    fmt.Println( srcIp[0] )
    //for i:= 0; i<4; i++ {
    //    fmt.Println()
    //}
}
