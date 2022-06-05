package main

import (
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	n = 400
)

type DistRequest struct {
	Name string
	Mode string
	Args []float64
	X    float64
}

type Distribution struct {
	distuv   DistuvDist
	xLo, xHi float64
	yLo, yHi float64
	discreet bool
}

func (d *Distribution) Prob(x float64) float64 {
	return d.distuv.Prob(x)
}

func (d *Distribution) CDF(x float64) float64 {
	return d.distuv.CDF(x)
}

func (d *Distribution) getPoints() []Point {
	var step float64
	var points []Point

	if d.discreet {
		step = 1.0
		points = make([]Point, int(d.xHi-d.xLo))
	} else {
		step = (d.xHi - d.xLo) / float64(n)
		points = make([]Point, n)
	}

	for i := range points {
		x := d.xLo + step*float64(i)
		y := d.Prob(x)
		points[i] = Point{x, y}

	}

	return points
}

func (d *Distribution) pointsUpTo(points []Point, x float64) []Point {
	upto := 0

	for i, point := range points {
		if point.X > x {
			upto = i
			break
		}
	}

	return points[:upto]
}

func (d *DistRequest) getData() GraphData {
	var data GraphData
	var dist Distribution

	switch d.Name {
	case "Normal":
		dist = Distribution{
			distuv:   distuv.Normal{Mu: d.Args[0], Sigma: d.Args[1]},
			xLo:      -4,
			xHi:      4,
			yLo:      0,
			yHi:      0.4,
			discreet: false,
		}

	case "Student's T":
		dist = Distribution{
			distuv:   distuv.StudentsT{Mu: d.Args[0], Sigma: d.Args[1], Nu: d.Args[2]},
			xLo:      -4,
			xHi:      4,
			yLo:      0,
			yHi:      0.4,
			discreet: false,
		}

	case "Chi-Squared":
		dist = Distribution{
			distuv:   distuv.ChiSquared{K: d.Args[0]},
			xLo:      0.000001,
			xHi:      40,
			yLo:      0,
			yHi:      0.25,
			discreet: false,
		}

	case "Gamma":
		dist = Distribution{
			distuv:   distuv.Gamma{Alpha: d.Args[0], Beta: d.Args[1]},
			xLo:      0.000001,
			xHi:      25,
			yLo:      0,
			yHi:      1,
			discreet: false,
		}

	case "Beta":
		dist = Distribution{
			distuv:   distuv.Beta{Alpha: d.Args[0], Beta: d.Args[1]},
			xLo:      0,
			xHi:      1,
			yLo:      0,
			yHi:      4,
			discreet: false,
		}

	case "Poisson":
		dist = Distribution{
			distuv:   distuv.Poisson{Lambda: d.Args[0]},
			xLo:      0,
			xHi:      25,
			yLo:      0,
			yHi:      0.4,
			discreet: true,
		}

	case "Binomial":
		dist = Distribution{
			distuv:   distuv.Binomial{N: d.Args[0], P: d.Args[1]},
			xLo:      0,
			xHi:      25,
			yLo:      0,
			yHi:      1,
			discreet: true,
		}

	case "Exponential":
		dist = Distribution{
			distuv:   distuv.Exponential{Rate: d.Args[0]},
			xLo:      0,
			xHi:      10,
			yLo:      0,
			yHi:      1,
			discreet: false,
		}
	}

	// set points for the curve
	data.Curve = dist.getPoints()

	// graph more data depending on what mode was selected
	switch d.Mode {
	case "Distribution":
		data.PDF = []Point{}
		data.CDF = dist.getPoints()
	case "PDF":
		data.PDF = []Point{{d.X, dist.Prob(d.X)}}
		data.CDF = []Point{}
	case "CDF":
		data.PDF = []Point{}
		data.CDF = dist.pointsUpTo(data.Curve, d.X)
	}

	return data
}
