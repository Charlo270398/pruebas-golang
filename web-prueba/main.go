package main

import (
	routes "./routes"
	log "./lib/logs"
	models "./models"
)

func main() {
	log.PrintLog("Servicio iniciado")
	models.CreateDB()
	routes.LoadRouter()
}
