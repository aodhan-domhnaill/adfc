package adfc

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	ics "github.com/arran4/golang-ical"
)

type Calendar struct {
	*ics.Calendar
	widget.BaseWidget

	layout *TimeAlignedLayout

	tabs *container.AppTabs
	box  *fyne.Container

	focusDate time.Time
}

func NewCalendar(ic *ics.Calendar, focus time.Time) *Calendar {
	layout := &TimeAlignedLayout{
		Start:    focus.Truncate(time.Hour * 24),
		Duration: time.Hour * 24,
	}
	c := &Calendar{
		Calendar: ic,
		layout:   layout,
		box:      container.New(layout),
	}
	c.ExtendBaseWidget(c)

	c.tabs = container.NewAppTabs(
		container.NewTabItem("Day", container.NewVScroll(c.box)),
		container.NewTabItem("Week", container.NewVScroll(c.box)),
	)

	c.tabs.OnSelected = func(ti *container.TabItem) {
		switch ti.Text {
		case "Day":
			layout.Duration = time.Hour * 24
			c.RefreshEvents()
		case "Week":
			layout.Duration = time.Hour * 24 * 7
			c.RefreshEvents()
		}
	}
	c.RefreshEvents()

	return c
}

func addTimeLines(c *fyne.Container, start time.Time, inc time.Duration, n uint8) {
	for i := 0; i < 24; i += 1 {
		c.Add(NewTimeAlignedObject(
			&canvas.Line{StrokeWidth: 1, StrokeColor: theme.ForegroundColor()},
			start.Add(time.Duration(i)*inc),
			0,
		))
	}
}

func (c *Calendar) RefreshEvents() {
	c.box.RemoveAll()

	vevents := c.Events()
	start, end := c.layout.Start, c.layout.Start.Add(c.layout.Duration)
	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		eod := d.AddDate(0, 0, 1)
		dayLayout := &TimeAlignedLayout{
			Start:    d,
			Duration: time.Hour * 24,
		}

		hourLines := container.New(dayLayout)
		addTimeLines(hourLines, d, time.Hour, 24)
		day := container.New(dayLayout)

		dayBox := container.NewBorder(
			&canvas.Text{
				Text:      d.Format(time.DateOnly),
				TextStyle: fyne.TextStyle{Bold: true},
				TextSize:  10,
			},
			nil,
			&canvas.Line{StrokeWidth: 1, StrokeColor: theme.ForegroundColor()},
			&canvas.Line{StrokeWidth: 1, StrokeColor: theme.ForegroundColor()},
			container.NewStack(hourLines, day),
		)

		for _, ve := range vevents {
			s, err := ve.GetStartAt()
			if err != nil {
				continue
			}

			e, err := ve.GetEndAt()
			if err != nil {
				continue
			}

			if (s.After(d) && s.Before(eod)) ||
				(e.After(d) && e.Before(eod)) {

				day.Add(NewTimeAlignedObject(
					NewEvent(ve), s, e.Sub(s),
				))
			}
		}

		c.box.Add(NewTimeAlignedObject(dayBox, d, time.Hour*24))
	}
}

func (c *Calendar) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.tabs)
}
