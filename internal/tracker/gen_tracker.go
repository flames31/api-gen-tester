package tracker

import (
	"time"

	"github.com/flames31/api-gen-tester/internal/log"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

func NewGenTracker() progress.Writer {
	log.L().Debug("created new tracker for generation")
	pw := progress.NewWriter()
	pw.SetTrackerLength(25)
	pw.SetUpdateFrequency(100 * time.Millisecond)
	pw.SetStyle(progress.StyleDefault)
	pw.SetAutoStop(false)
	pw.Style().Options.ErrorString = "ERROR"
	pw.Style().Options.Separator = ""
	pw.Style().Colors.Error = text.Colors{text.FgRed}

	passColor := text.Colors{text.FgGreen}
	pw.Style().Options.DoneString = passColor.Sprint("SUCCESS")

	return pw
}
