package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rcrespodev/Blogs/design/repository/api/v1"
	"github.com/rcrespodev/Blogs/design/repository/pkg/server/globalObjects"
	"log"
)

type Server struct {
	httpAddress string
	engine      *gin.Engine
}

func New(host string, port string) Server {
	server := Server{
		httpAddress: fmt.Sprintf("%s:%s", host, port),
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
	routes := v1.NewRoutes()
	for _, r := range routes.Routes {
		s.engine.Handle(r.HttpMethod(), r.RelativePath(), r.Handler())
	}
}
