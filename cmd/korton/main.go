package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/sirupsen/logrus"

	"github.com/keyopto/korton/internal/gui"
	"github.com/keyopto/korton/internal/logger"
)

func main() {
	logger.Logger = *logrus.New()
	logger.Logger.SetLevel(logrus.InfoLevel)

	myApp := app.New()
	gui.Setup(myApp)
}
