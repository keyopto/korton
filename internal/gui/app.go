package gui

import (
	"fyne.io/fyne/v2"
	"github.com/keyopto/korton/internal/gui/windows"
)

func Setup(myApp fyne.App) {
	windows.ShowAndRunMainWindow(myApp)
}
