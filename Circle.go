package main

import (
	gg "github.com/fogleman/gg"
)

const file = "aurea.png"

func main() {
	dc := gg.NewContext(5000, 5000)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(5)

	var x, y float64 = 2500, 2500

	fibonacciNums := GetNfibonacciElements(0, 1, 20)

	for i := 0; i < len(fibonacciNums); i++ {
		switch fibMod := i % 4; fibMod {
		case 0:
			dc.DrawArc(x, y, float64(fibonacciNums[i]), gg.Radians(0), gg.Radians(90))
			if i+1 != len(fibonacciNums) {
				y = y - (float64(fibonacciNums[i+1]) - float64(fibonacciNums[i]))
			}
		case 1:
			dc.DrawArc(x, y, float64(fibonacciNums[i]), gg.Radians(90), gg.Radians(180))
			if i+1 != len(fibonacciNums) {
				x = x + (float64(fibonacciNums[i+1]) - float64(fibonacciNums[i]))
			}
		case 2:

			dc.DrawArc(x, y, float64(fibonacciNums[i]), gg.Radians(180), gg.Radians(270))
			if i+1 != len(fibonacciNums) {
				y = y + (float64(fibonacciNums[i+1]) - float64(fibonacciNums[i]))
			}

		case 3:

			dc.DrawArc(x, y, float64(fibonacciNums[i]), gg.Radians(270), gg.Radians(360))
			if i+1 != len(fibonacciNums) {
				x = x - (float64(fibonacciNums[i+1]) - float64(fibonacciNums[i]))
			}

		}
	}

	dc.Stroke()
	dc.SavePNG("aurea.png")
}
