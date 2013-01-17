package main

import (
    "fmt"
    "strings"
    "strconv"
)

func computeDecIntPart(fp float64) int64 {
    s := fmt.Sprintf("%f", fp)
    s = s[2:]                      // strip off "0."
    s = strings.TrimRight(s, "0")  // strip off trailing zeros
    i, _ := strconv.ParseInt(s, 10, 64)
    return i
}

func ReadableFloat64(num float64, sep string) string {

    var (
        resultArr []string
        newFloat, newDecPart float64
        newIntPart, decIntPart int64
    )

    if num < 1000 {
        return fmt.Sprintf("%f -> %s", num,strings.TrimRight(fmt.Sprintf("%f", num), ".0"))
    }

    hasDecimal := false
    isNegative := false

    if num < 0 {
        isNegative = true
        num = num * -1
    }

    intPart := int64(num)
    decPart := num - float64(intPart)

    if decPart > 0 {
        hasDecimal = true
    }

    newFloat = float64(intPart) / 1000.0
    for {
        newIntPart = int64(newFloat)
        newDecPart = newFloat - float64(newIntPart)

        decIntPart = computeDecIntPart(newDecPart)
        if hasDecimal {
            resultArr = append(resultArr, fmt.Sprintf("%f", float64(decIntPart) + decPart))
            hasDecimal = false
        } else {
            resultArr = append(resultArr, fmt.Sprintf("%d", decIntPart))
        }

        if newIntPart < 1000 {
            resultArr = append(resultArr, fmt.Sprintf("%d", newIntPart))
            break
        }
        newFloat = float64(newIntPart) / 1000.0
    }

    if isNegative {
        lastElem := resultArr[len(resultArr)-1]
        resultArr[len(resultArr)-1] = fmt.Sprintf("-%s", lastElem)
    }

    // reverse order
    for i, j := 0, len(resultArr)-1; i < j; i, j = i+1, j-1 {
            resultArr[i], resultArr[j] = resultArr[j], resultArr[i]
    }
    retStr := strings.TrimRight(strings.Join(resultArr, sep), "0")

    return fmt.Sprintf("%f -> %s", num, retStr)

}

func main() {
    fmt.Println(ReadableFloat64(45, ","))
    fmt.Println(ReadableFloat64(123456789, ","))
    fmt.Println(ReadableFloat64(3456789.789, ","))
    fmt.Println(ReadableFloat64(-0123456789.789, ","))
    fmt.Println(ReadableFloat64(-10123456789.789, ","))
    fmt.Println(ReadableFloat64(00987654321.12345, ","))
    fmt.Println(ReadableFloat64(9123456789.0789, ",")) // rounding error .078899
}
