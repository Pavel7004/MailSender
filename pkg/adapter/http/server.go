package http

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Pavel7004/MailSender/pkg/adapter/http/v1"
	"github.com/Pavel7004/MailSender/pkg/infra/config"
)

// @title           MailSender API
// @version         0.1
// @description     This is an API for mail subscription service

// @contact.name   Kovalev Pavel
// @contact.email  kovalev5690@gmail.com

// @license.name   GPL-3.0
// @license.url    https://www.gnu.org/licenses/gpl-3.0.html

// @host      localhost:8080

type Server struct {
	router    *gin.Engine
	isRunning bool

	v1 *v1.Handler
}

func New(cfg *config.Config) *Server {
	server := new(Server)

	server.router = gin.New()
	server.isRunning = false
	server.v1 = v1.New(cfg)

	server.prepareRouter()

	return server
}

func (s *Server) Run() error {
	s.isRunning = true
	return s.router.Run(":8080")
}

func (s *Server) prepareRouter() {
}
