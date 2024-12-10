package general

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type LabelOptions struct {
	Color color.Color
	Size  float32
}

type LabelOptionFunction func(*LabelOptions)

func defaultLabelOptions() LabelOptions {
	return LabelOptions{
		Color: theme.Color(theme.ColorNameForeground),
		Size:  14,
	}
}

func LabelWithSize(size float32) LabelOptionFunction {
	return func(options *LabelOptions) {
		options.Size = size
	}
}

func LabelWithColor(colorArg color.Color) LabelOptionFunction {
	return func(options *LabelOptions) {
		options.Color = colorArg
	}
}

func CreateLabel(label string, userOptions ...LabelOptionFunction) *canvas.Text {
	options := defaultLabelOptions()

	for _, option := range userOptions {
		option(&options)
	}

	myLabel := canvas.NewText(label, options.Color)

	myLabel.TextSize = options.Size

	return myLabel
}
