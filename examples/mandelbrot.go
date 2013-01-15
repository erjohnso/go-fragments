package main

import (
    "fmt"
    "math/cmplx"
    "strings"
)

func toInfinity(cNum complex128, iterations int) bool {
    z := cNum
    for j := 0; j < iterations; j++ {
        z = cmplx.Pow(z, 2.0) + cNum
        if cmplx.Abs(z) > 4 { return true }
    }
    return false
}

func mandelbrot(xmin, xmax, ymin, ymax, xstep, ystep float64, iterations int) {
    var (
        b bool
        x, y float64
        z complex128
    )
    row := make([]string, int(xstep)+1)
    for yc := 0.0; yc < ystep; yc++ {
        y = yc*(ymax-ymin)/ystep + ymin
        row = make([]string, int(xstep)+1)
        for xc := 0.0; xc < xstep; xc++ {
            x = xc*(xmax-xmin)/xstep + xmin
            z = complex(x, y)
            b = toInfinity(z, iterations)
            if b {
                row = append(row, " ")
            } else {
                row = append(row, "*")
            }
        }
        fmt.Println(strings.Join(row, ""))
    }
}

func main() {
    mandelbrot(-2.0, 1.0, -1.0, 1.0, 64, 32, 1000)
}
