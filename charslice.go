package main

import "fmt"

func main() {
    var s1 string="GO GO GO"
    var s2 string="KOREA"
    var s3 string="MOONJAEIN"
    fmt.Println(len(s1))
    fmt.Println(len(s2))
    fmt.Println(len(s3))
    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3)
    for i := 0; i < len(s1); i++ {
        fmt.Println(s1[i:i+1])
    }
    for i := 0; i < len(s2); i++ {
        fmt.Println(s2[i:i+1])
    }
    for i := 0; i < len(s3); i++ {
        fmt.Println(s3[i:i+1])
    }
}
