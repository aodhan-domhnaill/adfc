package adfc

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TimeAlignedLayout struct {
	Start    time.Time
	Duration time.Duration
}

type TimeAlignedObject struct {
	widget.BaseWidget

	obj      fyne.CanvasObject
	Start    time.Time
	Duration time.Duration
}

func (tl *TimeAlignedLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	duration := tl.Duration.Seconds()
	for _, obj := range objects {
		tlObj, ok := obj.(*TimeAlignedObject)
		if ok {
			pos := fyne.NewPos(0,
				containerSize.Height*float32(tlObj.Start.Sub(tl.Start).Seconds()/duration))
			if pos.Y < 0 {
				pos.Y = 0
			}

			size := fyne.NewSize(containerSize.Width,
				containerSize.Height*float32(tlObj.Duration.Seconds()/duration))
			if size.Height > containerSize.Height {
				size.Height = containerSize.Height
			}

			obj.Move(pos)
			obj.Resize(size)
		} else {
			obj.Hide()
		}
	}
}

func (tl *TimeAlignedLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 800)
}

func NewTimeAlignedObject(obj fyne.CanvasObject, start time.Time, dur time.Duration) *TimeAlignedObject {
	tlObj := &TimeAlignedObject{
		obj:      obj,
		Start:    start,
		Duration: dur,
	}
	tlObj.ExtendBaseWidget(tlObj)

	return tlObj
}

func (tlObj *TimeAlignedObject) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(tlObj.obj)
}
