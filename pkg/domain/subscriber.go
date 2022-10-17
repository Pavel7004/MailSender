package domain

import (
	"time"
)

type Subcriber struct {
	Username string
	Name     string
	Surname  string
	Birthday time.Time
}

type AddSubscriberRequest struct {
	Username string
	Name     string
	Surname  string
	Birthday time.Time
}
