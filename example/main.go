package main

import (
	"os"

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
	window.SetContent(adfc.NewCalendar(cal))
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
