package main

import (
	"flag"
	"log"
	linker "ozonTask"
	"ozonTask/internal/preset"
	"ozonTask/pkg/handler"
	"ozonTask/pkg/link"
)

func main() {
	mode := flag.String("mode", "memory", "")
	var storage link.LinkStorage
	if *mode == "memory" {
		storage = preset.InitLinkMemory()
	} else if *mode == "db" {
		storage = preset.InitLinkSQL()
	} else {
		log.Fatalf("frong mode error")
	}
	linkHandler := handler.Handler{
		Repo: storage,
	}
	srv := new(linker.Server)

	_ = srv.Run("8080", linkHandler.InitRoutes())
}
