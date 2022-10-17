package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/Pavel7004/MailSender/pkg/domain"
	"github.com/Pavel7004/MailSender/pkg/infra/config"
)

type Handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) SendError(c *gin.Context, err error) {
	if e, ok := err.(*domain.Error); ok {
		c.JSON(e.CodeHTTP, e)
	} else {
		c.JSON(500, &domain.Error{
			Code:    "unknown_error",
			Message: err.Error(),
		})
	}
}
