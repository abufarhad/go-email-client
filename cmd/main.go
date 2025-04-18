package main

import (
	"email-client/internal/domain/service"
	"email-client/internal/infra/logger"
	"email-client/internal/interface/controller"
	"email-client/internal/interface/persistence"
	"email-client/internal/interface/ui"
)

func main() {
	logger.InitLogger()
	store := persistence.NewFileStore("emails.json")
	service := service.NewEmailService(store)
	handler := controller.NewHandler(service)
	ui.StartApp(handler)
}
