package adfc

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type DayLayout struct {
	title     fyne.CanvasObject
	leftLine  fyne.CanvasObject
	rightLine fyne.CanvasObject
	hourLines []fyne.CanvasObject
	date      time.Time
}

func NewDay(date time.Time) *fyne.Container {
	dl := &DayLayout{
		title: &canvas.Text{
			Text:      date.Format(time.DateOnly),
			TextStyle: fyne.TextStyle{Bold: true},
			TextSize:  10,
		},
		leftLine:  &canvas.Line{StrokeWidth: 1, StrokeColor: color.White},
		rightLine: &canvas.Line{StrokeWidth: 1, StrokeColor: color.White},
		date:      date,
	}

	c := container.New(dl)

	c.Add(dl.title)
	c.Add(dl.leftLine)
	c.Add(dl.rightLine)

	hourLines := make([]fyne.CanvasObject, 24)
	for i := 0; i < len(hourLines); i += 1 {
		hourLines[i] = &canvas.Line{StrokeWidth: 1, StrokeColor: color.White}
		c.Add(hourLines[i])
	}
	dl.hourLines = hourLines

	return c
}

func (dl *DayLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	titleHeight := float32(20)
	height := containerSize.Height - titleHeight
	width := containerSize.Width

	dl.title.Move(fyne.NewPos(0, 0))
	dl.title.Resize(fyne.NewSize(width, titleHeight))

	dl.leftLine.Move(fyne.NewPos(0, titleHeight))
	dl.leftLine.Resize(fyne.NewSize(0, height))

	dl.rightLine.Move(fyne.NewPos(width, titleHeight))
	dl.rightLine.Resize(fyne.NewSize(0, height))

	for i, hl := range dl.hourLines {
		hl.Move(fyne.NewPos(0, titleHeight+(height*float32(i))/24))
		hl.Resize(fyne.NewSize(width, 0))
	}

	dayStart := dl.date.Unix()
	daySec := float32((24 * time.Hour).Seconds())
	for _, obj := range objects {
		event, ok := obj.(*fyne.Container)
		if ok {
			eventLayout, ok := event.Layout.(EventLayout)
			if ok {
				start := eventLayout.Start()
				end := eventLayout.End()

				pos := fyne.NewPos(5, height*float32(start.Unix()-dayStart)/daySec)
				if pos.Y < 0 {
					pos.Y = 0
				}
				pos.Y += titleHeight

				size := fyne.NewSize(width-5, height*float32(end.Unix()-start.Unix())/daySec)
				if size.Height > height {
					size.Height = height
				}

				event.Move(pos)
				event.Resize(size)
			}
		}
	}
}

func (dl *DayLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 800)
}
