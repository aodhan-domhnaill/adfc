package main

import (
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/aodhan-domhnaill/adfc"
	ics "github.com/arran4/golang-ical"
)

func main() {
	f, err := os.Open("./events.ics")
	if err != nil {
		panic(err)
	}

	cal, err := ics.ParseCalendar(f)
	if err != nil {
		panic(err)
	}

	name := "Test"
	window := app.NewWithID(name).NewWindow(name)
	window.SetContent(adfc.NewCalendar(cal, time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC), color.RGBA{R: 255, A: 255}))
	window.Resize(fyne.NewSize(400, 800))
	window.ShowAndRun()
}
