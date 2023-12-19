package adfc

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type LayoutMode uint8

const (
	VerticalMode   LayoutMode = 0
	HorizontalMode LayoutMode = 1
)

type TimeAlignedLayout struct {
	Start    time.Time
	Duration time.Duration
	Mode     LayoutMode // Default Vertical
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
			posFactor := float32(tlObj.Start.Sub(tl.Start).Seconds() / duration)
			sizeFactor := float32(tlObj.Duration.Seconds() / duration)

			if posFactor < 0 {
				posFactor = 0
			}
			if sizeFactor > 1 {
				sizeFactor = 1
			}

			if tl.Mode == VerticalMode {
				obj.Move(fyne.NewPos(0, containerSize.Height*posFactor))
				obj.Resize(fyne.NewSize(containerSize.Width, containerSize.Height*sizeFactor))
			} else {
				obj.Move(fyne.NewPos(containerSize.Width*posFactor, 0))
				obj.Resize(fyne.NewSize(containerSize.Width*sizeFactor, containerSize.Height))
			}
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
