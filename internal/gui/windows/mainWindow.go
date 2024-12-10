package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/keyopto/korton/internal/gui/widgets/general"
)

func ShowAndRunMainWindow(myApp fyne.App) {
	window := myApp.NewWindow("Korton, the note taker")

	window.SetContent(container.NewVBox(
		container.NewCenter(
			general.CreateLabel(
				"Welcome, this is korton, the RPG note taker!",
				general.LabelWithSize(30),
			),
		),
		general.Button("Quit", func() {
			myApp.Quit()
		}),
	))

	window.Resize(fyne.NewSize(400, 200))
	window.ShowAndRun()
}
