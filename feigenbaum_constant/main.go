package main

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func logisticMap(r, x float64) float64 {
	return r * x * (1 - x)
}

func makeXYs(rValues, xValues []float64) plotter.XYs {
	if len(rValues) != len(xValues) {
		log.Fatal("rValues and xValues must have the same length")
	}
	xys := make(plotter.XYs, len(rValues))
	for i := range rValues {
		xys[i].X = rValues[i]
		xys[i].Y = xValues[i]
	}
	return xys
}

func main() {
	const iterations = 1000
	const last = 100
	const rMin = 2.4
	const rMax = 4.0
	const steps = 10000

	rValues := make([]float64, 0, steps*last)
	xValues := make([]float64, 0, steps*last)

	for r := rMin; r <= rMax; r += (rMax - rMin) / steps {
		x := 0.5
		for i := 0; i < iterations+last; i++ {
			x = logisticMap(r, x)
			if i >= iterations {
				rValues = append(rValues, r)
				xValues = append(xValues, x)
			}
		}
	}

	p := plot.New()

	p.Title.Text = "Bifurcation diagram"
	p.X.Label.Text = "r"
	p.Y.Label.Text = "x"

	scatter, err := plotter.NewScatter(makeXYs(rValues, xValues))
	if err != nil {
		log.Panic(err)
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	p.Add(scatter)

	if err := p.Save(10*vg.Inch, 4*vg.Inch, "bifurcation_diagram.png"); err != nil {
		log.Panic(err)
	}
}
