package adfc

import (
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type testLayout struct {
}

func (tl *testLayout) Start() time.Time {
	return time.Date(2023, 1, 1, 5, 0, 0, 0, time.UTC)
}

func (tl *testLayout) End() time.Time {
	return time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC)
}

func (tl *testLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
}

func (tl *testLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

func TestEmptyLayout(t *testing.T) {
	day := NewDay(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	dayLayout, ok := day.Layout.(*DayLayout)
	if !ok {
		t.Error("day has wrong layout")
	}

	event := container.New(&testLayout{})

	dayLayout.Layout(
		[]fyne.CanvasObject{
			event,
		},
		fyne.NewSize(100, 2420),
	)

	if len(dayLayout.hourLines) != 24 {
		t.Error("wrong number of hour lines")
	}
	for i, hl := range dayLayout.hourLines {
		if hl.Position().Y != float32(20+100*i) {
			t.Error("hour line positions wrong", hl.Position())
		}
	}

	if event.Position().Y != float32(520) {
		t.Error("event is misaligned")
	}
	if event.Size().Height != float32(100) {
		t.Error("event is missized")
	}
}
