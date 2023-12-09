package adfc

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type DayLayout struct {
}

func NewDay(date time.Time) *fyne.Container {
	c := container.New(layout.NewVBoxLayout())
	c.Add(&canvas.Text{
		Text:      date.Format(time.DateOnly),
		TextStyle: fyne.TextStyle{Bold: true},
		TextSize:  10,
	})
	return c
}

func (dl *DayLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
}

func (dl *DayLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 800)
}
