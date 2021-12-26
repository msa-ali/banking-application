package main

import (
	"github.com/Altamashattari/banking-application/app"
	"github.com/Altamashattari/banking-application/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
