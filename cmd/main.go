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
	mode := flag.String("mode", "db", "choose mode (memory/db)")
	var storage link.LinkStorage
	if *mode == "memory" {
		log.Println("used in-memory storage")
		storage = preset.InitLinkMemory()
	} else if *mode == "db" {
		log.Println("used database storage")
		storage = preset.InitLinkSQL()
	} else {
		log.Fatalf("wrong mode error")
	}
	linkHandler := handler.Handler{
		Repo: storage,
	}
	srv := new(linker.Server)

	_ = srv.Run("8080", linkHandler.InitRoutes())
}
