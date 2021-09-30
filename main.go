package main

import (
	"github.com/JyotsnaGorle/banking/app"
	"github.com/JyotsnaGorle/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
