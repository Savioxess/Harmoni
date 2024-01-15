package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var chordsInScale [7]string
var chordsDisplayText *widget.RichText

func CreateChordDisplay() *fyne.Container {
  getChordsInScale()

  chordsDisplayText = widget.NewRichTextFromMarkdown(fmt.Sprintf("## %s  %s  %s  %s  %s  %s  %s", chordsInScale[0], chordsInScale[1], chordsInScale[2], chordsInScale[3], chordsInScale[4], chordsInScale[5], chordsInScale[6]))
  
  chordsContainer := container.NewCenter(chordsDisplayText)

  return chordsContainer
}

func RefreshChords() {
  getChordsInScale()

  chordsDisplayText.ParseMarkdown(fmt.Sprintf("## %s  %s  %s  %s  %s  %s  %s", chordsInScale[0], chordsInScale[1], chordsInScale[2], chordsInScale[3], chordsInScale[4], chordsInScale[5], chordsInScale[6])) 

  chordsDisplayText.Refresh()
}

func getChordsInScale() {
  chordsForIonian := [...]string{"", "m", "m", "", "", "m", "dim"} 

  indexOfCurrentScale := getIndexOfCurrentScale()

  for index := 0; index < 7; index++ {
    chordsInScale[index] = fmt.Sprint(scale[index], chordsForIonian[getStepIndex(index+indexOfCurrentScale)])
  }
}
