package adfc

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	ics "github.com/arran4/golang-ical"
)

type eventLayout struct {
	e *ics.VEvent
}

func property(e *ics.VEvent, p ics.ComponentProperty) *canvas.Text {
	result := canvas.Text{
		Text:      "",
		TextStyle: fyne.TextStyle{},
		TextSize:  10,
		Color:     theme.ForegroundColor(),
	}

	r := e.GetProperty(p)
	if r != nil {
		result.Text = r.Value
	}
	return &result
}

func NewEvent(e *ics.VEvent) fyne.CanvasObject {
	return container.New(
		&eventLayout{e},
		&canvas.Rectangle{
			FillColor: theme.ShadowColor(),
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

func (el *eventLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	for _, obj := range objects {
		switch obj.(type) {
		case *canvas.Rectangle, *fyne.Container:
			obj.Move(fyne.NewPos(0, 0))
			obj.Resize(containerSize)
		}
	}
}

func (el *eventLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}
