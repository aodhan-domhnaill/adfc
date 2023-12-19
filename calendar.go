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

	layout *CalendarLayout

	tabs *container.AppTabs
	box  *fyne.Container

	focusDate time.Time
}

type CalendarMode int

const (
	DayMode   CalendarMode = 0
	WeekMode               = 1
	MonthMode              = 2
)

type CalendarLayout struct {
	FocusDate time.Time
	Mode      CalendarMode
}

func NewCalendar(ic *ics.Calendar, focus time.Time) *Calendar {
	layout := &CalendarLayout{
		Mode: DayMode,
	}
	c := &Calendar{
		Calendar: ic,
		layout:   layout,
		box:      container.New(layout),
	}
	c.ExtendBaseWidget(c)

	c.SetFocusDate(focus)

	c.tabs = container.NewAppTabs(
		container.NewTabItem("Day", container.NewVScroll(c.box)),
		container.NewTabItem("Week", container.NewVScroll(c.box)),
	)

	c.tabs.OnSelected = func(ti *container.TabItem) {
		switch ti.Text {
		case "Day":
			c.SetMode(DayMode)
		case "Week":
			c.SetMode(WeekMode)
		}
	}

	return c
}

func (c *Calendar) RefreshEvents() {
	c.box.RemoveAll()

	vevents := c.Events()
	start, end := c.layout.TimeRange()
	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		eod := d.AddDate(0, 0, 1)
		dayLayout := &TimeAlignedLayout{
			Start:    d,
			Duration: time.Hour * 24,
		}

		hourLines := container.New(dayLayout)
		day := container.New(dayLayout)
		for i := 0; i < 24; i += 1 {
			hourLines.Add(NewTimeAlignedObject(
				&canvas.Line{StrokeWidth: 1, StrokeColor: theme.ForegroundColor()},
				d.Add(time.Duration(i)*time.Hour),
				0,
			))
		}

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

		c.box.Add(dayBox)
	}
}

func (c *Calendar) SetFocusDate(t time.Time) {
	c.layout.FocusDate = t.Truncate(24 * time.Hour)
	c.RefreshEvents()
}

func (c *Calendar) SetMode(mode CalendarMode) {
	c.layout.Mode = mode
	c.RefreshEvents()
}

func (cl *CalendarLayout) TimeRange() (start time.Time, end time.Time) {
	switch cl.Mode {
	case DayMode:
		start = cl.FocusDate
		end = start.AddDate(0, 0, 1)
	case WeekMode:
		dow := int(cl.FocusDate.Weekday())
		start = cl.FocusDate.AddDate(0, 0, -1*dow)
		end = start.AddDate(0, 0, 7)
	}
	return
}

func (c *Calendar) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.tabs)
}

func (cl *CalendarLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var objSize fyne.Size
	switch cl.Mode {
	case DayMode:
		objSize = containerSize
	case WeekMode:
		objSize = fyne.NewSize(containerSize.Width/7, containerSize.Height)
	}

	// TODO: Add a Day object with the timeline
	pos := fyne.NewPos(0, 0)
	for _, obj := range objects {
		obj.Resize(objSize)
		obj.Move(pos)

		pos.X += objSize.Width
		if pos.X >= containerSize.Width {
			pos.X = 0
			pos.Y += objSize.Height
		}
	}
}

func (cl *CalendarLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 800)
}
