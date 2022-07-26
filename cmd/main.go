package main

import (
	linker "ozonTask"
	"ozonTask/internal/preset"
	"ozonTask/pkg/handler"
)

func main() {

	storage := preset.GetStorage()

	linkHandler := handler.Handler{
		Repo: storage,
	}

	srv := new(linker.Server)

	_ = srv.Run("8080", linkHandler.InitRoutes())
}
