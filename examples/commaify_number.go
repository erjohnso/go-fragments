package main

import (
    "fmt"
    "strings"
)

func CommaifyFloat64(fp float64) string {
    var (
        isNegative bool
        strfp string
        parts []string
        result []string
        l int
    )
    if fp < 0 {
        isNegative = true
        fp = fp * -1
    }
    // convert to string and trim off leading/trailing zeros
    strfp = strings.Trim(fmt.Sprintf("%f", fp), "0")
    // split on decimal point
    parts = strings.Split(strfp, ".")
    if len(parts) == 2 && len(parts[1]) > 0 {
        result = append(result, parts[1], ".")
    }
    // is number >= 100?
    if len(parts[0]) >= 3 {
        tmp := parts[0]
        for len(tmp) >= 3 {
            l = len(tmp)
            result = append(result, tmp[(l-3):l])
            tmp = tmp[0:(l-3)]
        }
        if len(tmp) > 0 {
            result = append(result, tmp)
        }
    } else {
        result = append(result, parts[0])
    }
    // reverse list
    for i, j := 0, len(result) - 1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    if isNegative {
        return strings.Replace(
            fmt.Sprintf("-%s", strings.Join(result, ",")), ",.,", ".", 1)
    }
    return strings.Replace(
        fmt.Sprintf("%s", strings.Join(result, ",")), ",.,", ".", 1)
}

func main() {
    fmt.Printf("1. %s\n", CommaifyFloat64(45))
    fmt.Printf("2. %s\n", CommaifyFloat64(123456789))
    fmt.Printf("3. %s\n", CommaifyFloat64(34567890.0123))
    fmt.Printf("4. %s\n", CommaifyFloat64(-0123456789.789))
    fmt.Printf("5. %s\n", CommaifyFloat64(-10123456789.789))
    fmt.Printf("6. %s\n", CommaifyFloat64(00987654321.1234500))
    fmt.Printf("7. %s (rounding error)\n", CommaifyFloat64(9123456789.0789))
}
