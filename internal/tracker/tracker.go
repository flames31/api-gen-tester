package tracker

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

func NewTracker() progress.Writer {
	pw := progress.NewWriter()
	pw.SetTrackerLength(25)
	pw.SetUpdateFrequency(100 * time.Millisecond)
	pw.SetStyle(progress.StyleDefault)
	pw.SetAutoStop(false)
	pw.Style().Options.ErrorString = "FAIL"
	pw.Style().Options.Separator = ""
	pw.Style().Colors.Error = text.Colors{text.FgRed}

	passColor := text.Colors{text.FgGreen}
	pw.Style().Options.DoneString = passColor.Sprint("PASS")

	return pw
}
