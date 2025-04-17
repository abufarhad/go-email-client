package main

import (
	"email-client/internal/backend"
	"email-client/internal/controller"
	"email-client/internal/ui"
)

func main() {
	store := backend.NewStore()
	handler := controller.NewHandler(store)
	ui.StartApp(handler)
}
