package adfc

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	ics "github.com/arran4/golang-ical"
)

type EventLayout struct {
	e *ics.VEvent
}

func property(e *ics.VEvent, p ics.ComponentProperty) *canvas.Text {
	result := canvas.Text{
		Text:      "",
		TextStyle: fyne.TextStyle{},
		TextSize:  10,
		Color:     color.Black,
	}

	r := e.GetProperty(p)
	if r != nil {
		result.Text = r.Value
	}
	return &result
}

func NewEvent(e *ics.VEvent) fyne.CanvasObject {
	return container.New(
		&EventLayout{e},
		&canvas.Rectangle{
			FillColor: color.White,
		},
		container.NewVBox(
			property(e, ics.ComponentPropertySummary),
			property(e, ics.ComponentPropertyDtStart),
			property(e, ics.ComponentPropertyDtStart),
			property(e, ics.ComponentPropertyDescription),
			property(e, ics.ComponentPropertyLocation),
		),
	)
}

func (el *EventLayout) Start() time.Time {
	start, err := el.e.GetStartAt()
	if err != nil {
		fyne.LogError("error in DayLayout", err)
	}
	return start
}

func (el *EventLayout) End() time.Time {
	end, err := el.e.GetEndAt()
	if err != nil {
		fyne.LogError("error in DayLayout", err)
	}
	return end
}

func (el *EventLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	for _, obj := range objects {
		switch obj.(type) {
		case *canvas.Rectangle, *fyne.Container:
			obj.Move(fyne.NewPos(0, 0))
			obj.Resize(containerSize)
		}
	}
}

func (el *EventLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}
