package adfc

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	ics "github.com/arran4/golang-ical"
)

type Calendar struct {
	*ics.Calendar
	widget.BaseWidget

	tabs   *container.AppTabs
	days   map[time.Time]day
	agenda fyne.Container
}

type day struct {
	date   time.Time
	events []*ics.VEvent
	vbox   *fyne.Container
}

func NewCalendar(ic *ics.Calendar) *Calendar {
	c := &Calendar{
		Calendar: ic,
		days:     map[time.Time]day{},
	}
	c.ExtendBaseWidget(c)

	c.arrangeDays()

	c.tabs = container.NewAppTabs(
		container.NewTabItem("Agenda", container.NewVScroll(&c.agenda)),
		container.NewTabItem("Week", container.NewVScroll(&c.agenda)),
	)

	c.agenda.Layout = layout.NewVBoxLayout()
	c.tabs.OnSelected = func(ti *container.TabItem) {
		switch ti.Text {
		case "Agenda":
			c.agenda.Layout = layout.NewVBoxLayout()
		case "Week":
			c.agenda.Layout = layout.NewHBoxLayout()
		}
	}

	return c
}

func (c *Calendar) arrangeDays() {
	oneday := 24 * time.Hour
	for t, events := range c.PartitionEvents(&oneday) {
		d := day{
			date:   t,
			events: events,
			vbox:   container.NewVBox(),
		}

		for _, e := range events {
			ec := NewEvent(e)
			d.vbox.Add(ec)
		}

		c.agenda.Add(d.vbox)
		c.days[t] = d
	}
}

func (c *Calendar) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.tabs)
}

func (c *Calendar) PartitionEvents(duration *time.Duration) map[time.Time][]*ics.VEvent {
	events := map[time.Time][]*ics.VEvent{}

	for _, e := range c.Calendar.Events() {
		t, err := e.GetStartAt()
		if err != nil {
			fyne.LogError("unable to get start at", err)
		} else {
			tr := t.Truncate(*duration)
			events[tr] = append(events[tr], e)
		}
	}

	return events
}
