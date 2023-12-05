package adfc

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	ics "github.com/arran4/golang-ical"
)

type Event struct {
	*ics.VEvent
	widget.BaseWidget

	card *widget.Card

	startAt, endAt *time.Time
}

func NewEvent(e *ics.VEvent) *Event {
	event := &Event{
		VEvent: e,
		card:   &widget.Card{},
	}
	event.ExtendBaseWidget(event)

	startTime, err := e.GetStartAt()
	if err == nil {
		event.startAt = &startTime
	}
	endTime, err := e.GetEndAt()
	if err == nil {
		event.endAt = &endTime
	}

	event.card.Subtitle = startTime.String() + " - " + endTime.String()

	summary := e.GetProperty(ics.ComponentPropertySummary)
	if summary != nil {
		event.card.Title = summary.Value
	}

	description := e.GetProperty(ics.ComponentPropertyDescription)
	location := e.GetProperty(ics.ComponentPropertyLocation)
	if location != nil && description != nil {
		event.card.Content = widget.NewRichTextFromMarkdown(location.Value)
	}

	return event
}

func (e *Event) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.card)
}
