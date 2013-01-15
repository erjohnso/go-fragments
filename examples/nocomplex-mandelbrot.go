package main

import (
    "fmt"
    "strings"
)

func toInfinity(realPart, imgPart, iterations float64) bool {
    zr := 0.0
    zi := 0.0
    for j := 0.0; j < iterations; j++ {
        zr, zi = zr*zr - zi*zi + realPart, 2*zr*zi + imgPart
        if zr*zr + zi*zi > 4 {
            return true
        }
    }
    return false
}

func mandelbrot(xmin, xmax, xstep, ymin, ymax, ystep, iterations float64) {
    row := make([]string, int(xstep)+1)
    b := false
    y := 0.0
    x := 0.0
    for yc := 0.0; yc < ystep; yc++ {
        y = yc*(ymax-ymin)/ystep + ymin
        row = make([]string, int(xstep)+1)
        for xc := 0.0; xc < xstep; xc++ {
            x = xc*(xmax-xmin)/xstep + xmin
            b = toInfinity(x, y, iterations)
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
    mandelbrot(-2.0, 1.0, 64, -1.0, 1.0, 32, 1000)
}
