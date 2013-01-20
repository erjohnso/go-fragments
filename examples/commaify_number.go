package main

import (
    "fmt"
    "strings"
)

func CommaifyFloat64(fp float64, sep string) string {
    var (
        isNegative bool
        strfp string
        parts []string
        result []string
    )
    if fp < 0 {
        isNegative = true
        fp = fp * -1
    }
    strfp = strings.Trim(fmt.Sprintf("%f", fp), "0")
    parts = strings.Split(strfp, ".")
    if len(parts) == 2 {
        result = append(result, parts[1])
    }
    for i := len(parts[1]); i > 3; i-=3 {
        result = append(result, parts[1][i:(i-3)])
    }
    if isNegative {
        result = append(result, "-")
    }
    fmt.Println(result)
    return ""
}

func main() {
    fmt.Println(CommaifyFloat64(45, ","))
    fmt.Println(CommaifyFloat64(123456789, ","))
    fmt.Println(CommaifyFloat64(3456789.789, ","))
    fmt.Println(CommaifyFloat64(-0123456789.789, ","))
    fmt.Println(CommaifyFloat64(-10123456789.789, ","))
    fmt.Println(CommaifyFloat64(00987654321.12345, ","))
    fmt.Println(CommaifyFloat64(9123456789.0789, ",")) // rounding error .078899
}
