package adfc

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	ics "github.com/arran4/golang-ical"
)

type Event struct {
	*ics.VEvent
	widget.BaseWidget

	startAt, endAt                 *canvas.Text
	summary, description, location *canvas.Text

	box *fyne.Container
}

func property(e *ics.VEvent, p ics.ComponentProperty) *canvas.Text {
	result := canvas.Text{
		Text:      "",
		TextStyle: fyne.TextStyle{},
		TextSize:  10,
	}

	r := e.GetProperty(p)
	if r != nil {
		result.Text = r.Value
	}
	return &result
}

func NewEvent(e *ics.VEvent) *Event {
	event := &Event{
		VEvent:      e,
		startAt:     property(e, ics.ComponentPropertyDtStart),
		endAt:       property(e, ics.ComponentPropertyDtStart),
		summary:     property(e, ics.ComponentPropertySummary),
		description: property(e, ics.ComponentPropertyDescription),
		location:    property(e, ics.ComponentPropertyLocation),
	}
	event.ExtendBaseWidget(event)

	event.summary.TextStyle = fyne.TextStyle{Bold: true}
	event.startAt.TextStyle = fyne.TextStyle{Monospace: true}
	event.description.TextStyle = fyne.TextStyle{Italic: true}

	event.box = container.NewVBox(
		event.summary,
		event.startAt,
		event.description,
		event.location,
	)

	return event
}

func (e *Event) Resize(size fyne.Size) {
	if size.Height < 200 {
		e.summary.TextSize = 20
	} else if size.Height < 400 {
		e.summary.TextSize = 30
	} else if size.Height < 800 {
		e.summary.TextSize = 40
	}
}

func (e *Event) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.box)
}
