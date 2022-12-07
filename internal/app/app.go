package app

import (
	"fmt"

	"github.com/MaxWxaM/linebot/internal/app/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	o      *config.HttpConfig
}

func NewServer(o *config.HttpConfig, r *gin.Engine) *Server {
	server := &Server{
		o:      o,
		router: r,
	}
	return server
}

func (s *Server) Start() error {
	port := fmt.Sprintf(":%v", s.o.Port)
	return s.router.Run(port)
}
