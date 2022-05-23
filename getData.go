package main

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	n = 400
)

type DistRequest struct {
	Name string
	Mode string
	Args []float64
}

func (d *DistRequest) getData() GraphData {
	var data GraphData

	switch d.Name {
	case "Normal":
		dist := distuv.Normal{
			Mu:    d.Args[0],
			Sigma: d.Args[1],
		}
		begin, end := dist.Mu-4.0*dist.Sigma, dist.Mu+4.0*dist.Sigma
		data.Points = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[2], dist.Prob(d.Args[2])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := dist.Mu-4.0*dist.Sigma, math.Min(d.Args[2], 4.0*dist.Sigma)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	case "Student's T":
		dist := distuv.StudentsT{
			Mu:    d.Args[0],
			Sigma: d.Args[1],
			Nu:    d.Args[2],
		}
		begin, end := dist.Mu-4.0*dist.Sigma, dist.Mu+4.0*dist.Sigma
		data.Points = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[3], dist.Prob(d.Args[3])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := dist.Mu-4.0*dist.Sigma, math.Min(d.Args[3], 4.0*dist.Sigma)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	case "Chi-Squared":
		dist := distuv.ChiSquared{
			K: d.Args[0],
		}
		begin, end := 0.000001, 25.0
		data.Points = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[1], dist.Prob(d.Args[1])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := 0.000001, math.Min(d.Args[1], 25.0)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	case "Gamma":
		dist := distuv.Gamma{
			Alpha: d.Args[0],
			Beta:  d.Args[1],
		}
		begin, end := 0.000001, 25.0
		data.Points = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[2], dist.Prob(d.Args[2])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := 0.000001, math.Min(d.Args[2], 25.0)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	case "Beta":
		dist := distuv.Beta{
			Alpha: d.Args[0],
			Beta:  d.Args[1],
		}
		begin, end := 0.0, 1.0
		data.Points = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[2], dist.Prob(d.Args[2])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := 0.0, math.Min(d.Args[2], 25.0)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	case "Poisson":
		dist := distuv.Poisson{
			Lambda: d.Args[0],
		}
		data.Points = make([]Point, 26)
		for i := range data.Points {
			data.Points[i] = Point{float64(i), dist.Prob(float64(i))}
		}

		switch d.Mode {
		case "Distribution":
			data.CDF = data.Points
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[1], dist.Prob(d.Args[1])}}
			data.CDF = []Point{}
		case "CDF":
			for i := range data.Points {
				if data.Points[i].X <= d.Args[1] {
					data.CDF = append(data.CDF, data.Points[i])
				}
			}
			data.PDF = []Point{}
		}

	case "Binomial":
		dist := distuv.Binomial{
			N: d.Args[0],
			P: d.Args[1],
		}
		data.Points = make([]Point, int(dist.N)+1)
		for i := range data.Points {
			data.Points[i] = Point{float64(i), dist.Prob(float64(i))}
		}

		switch d.Mode {
		case "Distribution":
			data.CDF = data.Points
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[2], dist.Prob(d.Args[2])}}
			data.CDF = []Point{}
		case "CDF":
			for i := range data.Points {
				if data.Points[i].X <= d.Args[2] {
					data.CDF = append(data.CDF, data.Points[i])
				}
			}
			data.PDF = []Point{}
		}

	case "Exponential":
		dist := distuv.Exponential{
			Rate: d.Args[0],
		}
		begin, end := 0.0, 25.0
		data.Points = curvePoints(dist, begin, end)

		switch d.Mode {
		case "Distribution":
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		case "PDF":
			data.PDF = []Point{{d.Args[1], dist.Prob(d.Args[1])}}
			data.CDF = []Point{}
		case "CDF":
			begin, end := 0.0, math.Min(d.Args[1], 25.0)
			data.CDF = curvePoints(dist, begin, end)
			data.PDF = []Point{}
		}

	}

	return data
}
