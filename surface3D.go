package main

import (
	"io"
	"log"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
)

var surfaceRangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genSurface3dData0() [][3]interface{} {
	data := make([][3]interface{}, 0)

	for i := -60; i < 60; i++ {
		y := float64(i) / 60
		for j := -60; j < 60; j++ {
			x := float64(j) / 60
			z := math.Sin(x*math.Pi) * math.Sin(y*math.Pi)
			data = append(data, [3]interface{}{x, y, z})
		}
	}
	return data
}

func genSurface3dData1() [][3]interface{} {
	data := make([][3]interface{}, 0)
	for i := -30; i < 30; i++ {
		y := float64(i) / 10
		for j := -30; j < 30; j++ {
			x := float64(j) / 10
			z := math.Sin(x*x+y*y) * x / math.Pi
			data = append(data, [3]interface{}{x, y, z})
		}
	}
	return data
}

func surface3DBase() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "surface3D-example",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        3,
			Min:        -3,
		}),
	)

	surface3d.AddZAxis("surface3d", genSurface3dData0())
	return surface3d
}

func surface3DRose() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "surface3D-Rose",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        3,
			Min:        -3,
		}),
	)

	surface3d.AddZAxis("surface3d", genSurface3dData1())
	return surface3d
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		surface3DBase(),
		surface3DRose(),
	)

	f, err := os.Create("surface3D.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))
}
