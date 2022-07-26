package linker

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{Addr: ":" + port,
		Handler: handler}
	log.Printf("start listen on %s", port)
	return s.httpServer.ListenAndServe()

}
