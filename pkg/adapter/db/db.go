package db

import (
	"context"

	"github.com/Pavel7004/MailSender/pkg/domain"
)

type DB interface {
	Close() error

	AddSubscriber(ctx context.Context, req *domain.AddSubscriberRequest) error
}
