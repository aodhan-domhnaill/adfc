package adfc

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	ics "github.com/arran4/golang-ical"
)

type Calendar struct {
	*ics.Calendar
}

func NewCalendar(ic *ics.Calendar, focus time.Time, fillColor color.Color) *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Day", container.NewVScroll(day(ic, focus, fillColor))),
		container.NewTabItem("Week", container.NewVScroll(week(ic, focus, fillColor))),
	)

	return tabs
}

func timelines(start time.Time, inc time.Duration, n uint8) (lines []fyne.CanvasObject) {
	for i := 0; i < 24; i += 1 {
		lines = append(lines,
			NewTimeAlignedObject(
				&canvas.Line{StrokeWidth: 1, StrokeColor: theme.ForegroundColor()},
				start.Add(time.Duration(i)*inc),
				0,
			),
		)
	}
	return
}

func inRange(c *ics.Calendar, start, end time.Time, fillColor color.Color) (events []fyne.CanvasObject) {
	for _, ve := range c.Events() {
		s, err := ve.GetStartAt()
		if err != nil {
			continue
		}

		e, err := ve.GetEndAt()
		if err != nil {
			continue
		}

		if (s.After(start) && s.Before(end)) ||
			(e.After(start) && e.Before(end)) {

			events = append(events, NewTimeAlignedObject(
				NewEvent(ve, fillColor), s, e.Sub(s),
			))
		}
	}
	return
}

func day(c *ics.Calendar, d time.Time, fillColor color.Color) *TimeAlignedObject {
	dayLayout := &TimeAlignedLayout{
		Start:    d,
		Duration: time.Hour * 24,
	}

	dayBox := container.NewVBox(
		&canvas.Text{
			Text:      d.Format(time.DateOnly),
			TextStyle: fyne.TextStyle{Bold: true},
			TextSize:  10,
		},
		container.NewStack(
			container.New(dayLayout, timelines(d, time.Hour, 24)...),
			container.New(dayLayout, inRange(c, d, d.AddDate(0, 0, 1), fillColor)...),
		),
	)
	return NewTimeAlignedObject(dayBox, d, time.Hour*24)
}

func week(c *ics.Calendar, d time.Time, fillColor color.Color) *TimeAlignedObject {
	weekLayout := &TimeAlignedLayout{
		Start:    d,
		Duration: time.Hour * 24 * 7,
		Mode:     HorizontalMode,
	}

	var days []fyne.CanvasObject
	for i := 0; i < 7; i += 1 {
		days = append(days, day(c, d.AddDate(0, 0, i), fillColor))
	}

	weekBox := container.NewVBox(
		container.NewStack(
			container.New(weekLayout, timelines(d, time.Hour*24, 7)...),
			container.New(weekLayout, days...),
		),
	)
	return NewTimeAlignedObject(weekBox, d, time.Hour*24*7)
}
