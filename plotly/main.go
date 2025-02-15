package main

import (
	"time"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Bar{
				X: types.DataArray([]time.Time{time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 4, 2, 0, 0, 0, 0, time.Local), time.Date(2022, 4, 3, 0, 0, 0, 0, time.Local)}),
				Y: types.DataArray([]float64{1, 2, 3}),
			},
			&grob.Bar{
				X: types.DataArray([]time.Time{time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2022, 4, 2, 0, 0, 0, 0, time.Local), time.Date(2022, 4, 3, 0, 0, 0, 0, time.Local)}),
				Y: types.DataArray([]float64{1, 3, 2}),
			},
		},
		Layout: &grob.Layout{
			Barmode: "stack",
			Title: &grob.LayoutTitle{
				Text: "Stacked Bar Chart",
			},
			Xaxis: &grob.LayoutXaxis{
				Tickformat: "%Y/%m/%d",
			},
		},
	}

	offline.ToHtml(fig, "bar.html")
	offline.Show(fig)
}
