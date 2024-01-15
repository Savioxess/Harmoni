package main

import (
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
  // "fyne.io/fyne/v2/container"
)

func main() {
  appInstance := app.New()
  window := appInstance.NewWindow("Harmoni")

  window.SetContent(widget.NewLabel("Find diatonic chords to any scale"))

  window.Show()
  appInstance.Run()
}
