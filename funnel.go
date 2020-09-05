package main

import (
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/components"
	"github.com/go-echarts/go-echarts/opts"
)

var dimensions = []string{"Visit", "Add", "Order", "Payment", "Deal"}

func genFunnelKvItems() []opts.FunnelData {
	items := make([]opts.FunnelData, 0)
	for i := 0; i < len(dimensions); i++ {
		items = append(items, opts.FunnelData{Name: dimensions[i], Value: rand.Intn(50)})
	}
	return items
}
func funnelBase() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Funnel-basic-example",
		}),
	)

	funnel.AddSeries("Analytics", genFunnelKvItems())
	return funnel
}

// TODO: check the different from echarts side
func funnelShowLabel() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Funnel-show-label",
		}),
	)

	funnel.AddSeries("Analytics", genFunnelKvItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:     true,
				Position: "left",
			},
		))
	return funnel
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		funnelBase(),
		funnelShowLabel(),
	)

	f, err := os.Create("funnel.html")
	if err != nil {
		log.Println(err)
	}
	_ = page.Render(io.MultiWriter(os.Stdout, f))

}
