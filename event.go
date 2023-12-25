package adfc

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	ics "github.com/arran4/golang-ical"
)

type eventLayout struct {
	s []string
}

func property(e *ics.VEvent, p ics.ComponentProperty) string {
	r := e.GetProperty(p)
	if r != nil {
		return r.Value
	}
	return ""
}

func NewEvent(e *ics.VEvent, fillColor color.Color) fyne.CanvasObject {
	t := []fyne.CanvasObject{
		&canvas.Text{
			TextStyle: fyne.TextStyle{Bold: true},
			TextSize:  14,
			Text:      property(e, ics.ComponentPropertySummary),
		},
		&canvas.Text{TextSize: 10, Text: property(e, ics.ComponentPropertyDtStart)},
		&canvas.Text{TextSize: 10, Text: property(e, ics.ComponentPropertyDtEnd)},
		&canvas.Text{TextSize: 10, Text: property(e, ics.ComponentPropertyDescription)},
		&canvas.Text{TextSize: 10, Text: property(e, ics.ComponentPropertyLocation)},
	}

	s := make([]string, len(t))
	for i, tt := range t {
		s[i] = tt.(*canvas.Text).Text
	}

	return container.NewStack(
		&canvas.Rectangle{
			FillColor:    fillColor,
			StrokeColor:  theme.InputBorderColor(),
			CornerRadius: 10,
		},
		container.New(&eventLayout{s}, t...),
	)
}

func fitToWidth(orig string, width float32, w *canvas.Text) string {
	text := orig
	size := fyne.MeasureText(text, w.TextSize, w.TextStyle)
	for size.Width > width {
		text = text[:len(text)-1]
		size = fyne.MeasureText(text+"...", w.TextSize, w.TextStyle)
	}
	if len(text) < len(orig) {
		return text + "..."
	}
	return orig
}

func (el *eventLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	y := float32(0)
	for i, obj := range objects {
		w := obj.(*canvas.Text)
		w.Text = fitToWidth(el.s[i], containerSize.Width, w)

		obj.Move(fyne.NewPos(0, y))
		y += obj.MinSize().Height
		if y > containerSize.Height {
			obj.Hide()
		}
	}
}

func (el *eventLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}
