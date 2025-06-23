package generate

import (
	"github.com/flames31/api-gen-tester/internal/tracker"
	"github.com/jedib0t/go-pretty/v6/progress"
)

func startTracker() (progress.Writer, *progress.Tracker) {
	pw := tracker.NewGenTracker()

	tr := &progress.Tracker{
		Message: "Generating test cases : ",
		Total:   100,
		Units:   progress.UnitsDefault,
	}
	pw.AppendTracker(tr)
	go pw.Render()

	tr.SetValue(20)

	return pw, tr
}
