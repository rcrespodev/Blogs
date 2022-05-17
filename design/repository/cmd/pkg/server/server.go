package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/server/globalObjects"
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/server/handlers"
	"log"
)

type Server struct {
	httpAddress string
	engine      *gin.Engine
}

func New(host string, port uint) Server {
	server := Server{
		httpAddress: fmt.Sprintf("%s:%d", host, port),
		engine:      gin.New(),
	}

	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	if err := globalObjects.New(); err != nil {
		return err
	}

	log.Println("server running on", s.httpAddress)
	return s.engine.Run(s.httpAddress)
}

func (s *Server) registerRoutes() {
	routes := handlers.NewRoutes()
	for _, r := range routes.Routes {
		s.engine.Handle(r.HttpMethod(), r.RelativePath(), r.Handler())
	}
}
