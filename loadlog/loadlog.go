package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/MetalBlueberry/go-plotly/offline"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {

	fileBuf, err := loadFile("logfile.log")
	if err != nil {
		panic(err)
	}

	regs := map[*regexp.Regexp][]string{
		regexp.MustCompile(`.*func1 Start`): []string{"func1"},
		regexp.MustCompile(`.*func1 End`):   []string{"func1"},
	}

	rDateTime := regexp.MustCompile(`\[(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}):(\d{3})\](.*)`)
	layout := "2006/01/02 15:04:05.000"
	loc, _ := time.LoadLocation("Local")

	colDateTimeStr := []string{}
	colDateTime := []time.Time{}
	colMethod := []string{}
	colMessage := []string{}

	now := time.Now()

	// df := dataframe.New(
	// 	series.New([]string{}, series.String, "DATETIME"),
	// 	series.New([]string{}, series.String, "METHOD"),
	// 	series.New([]string{}, series.String, "MESSAGE"),
	// )

	lines := strings.Split(fileBuf, "\r\n")
	fmt.Printf("file length: %v\n", len(lines))

	for _, line := range lines {
		for k, vv := range regs {
			if k.MatchString(line) {
				result := rDateTime.FindAllStringSubmatch(line, -1)
				strDateTime := fmt.Sprintf("%s.%s", result[0][1], result[0][2])
				dtDateTime, _ := time.ParseInLocation(layout, strDateTime, loc)

				for _, v := range vv {
					// df2 := dataframe.LoadRecords(
					// 	[][]string{
					// 		[]string{"DATETIME", "METHOD", "MESSAGE"},
					// 		[]string{strDateTime, v, result[0][3]},
					// 	},
					// )
					// df = df.RBind(df2)

					colDateTimeStr = append(colDateTimeStr, strDateTime)
					colDateTime = append(colDateTime, dtDateTime)
					colMethod = append(colMethod, v)
					colMessage = append(colMessage, result[0][3])
				}

				break
			}
		}
	}

	df := dataframe.New(
		series.New(colDateTimeStr, series.String, "DATETIME"),
		series.New(colMethod, series.String, "METHOD"),
		series.New(colMessage, series.String, "MESSAGE"),
	)

	fmt.Printf("time: %vms\n", time.Since(now).Milliseconds())
	// fmt.Println(df)

	csvFile, _ := os.Create("data.csv")
	defer csvFile.Close()

	df.WriteCSV(csvFile)

	// dfSel := df.Select([]string{"DATETIME"})
	// fmt.Println(dfSel)
	// fmt.Println(df.Col("DATETIME"))

	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter{
				Type: grob.TraceTypeScatter,
				X:    colDateTime,
				Y:    colMethod,
				Mode: grob.ScatterModeMarkers + "+" + grob.ScatterModeLines,
				Text: colMessage,
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure",
			},
		},
	}

	// offline.Show(fig)
	offline.ToHtml(fig, "graph.html")

}

func loadFile(path string) (string, error) {

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
