package src

import (
	"bytes"
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func generateTextReport(metrics *vegeta.Metrics) {

	var b bytes.Buffer
	textReporter := vegeta.NewTextReporter(metrics)
	err := textReporter(&b)
	os.Stdout.WriteString(b.String())

	check(err)
}

func generatePlotReport(results *vegeta.Results) {
	var b bytes.Buffer

	t := time.Now()
	loadTest := fmt.Sprintf("%d%02d%02dT%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	plotReporter := vegeta.NewPlotReporter("load-test-hsl-"+loadTest, results)
	err := plotReporter(&b)
	check(err)

	f, err := os.Create("report/load-test-hsl-" + loadTest + ".html")
	check(err)
	defer f.Close()
	_, err = f.WriteString(b.String())
	check(err)

}
