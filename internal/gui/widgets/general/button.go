package general

import "fyne.io/fyne/v2/widget"

func Button(label string, action func()) *widget.Button {
	return widget.NewButton(label, action)
}
