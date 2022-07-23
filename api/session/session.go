package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID
	ExpiredAt time.Time
}

func(s *Session) IsExpired() bool {
	return s.ExpiredAt.Before(time.Now())
}
