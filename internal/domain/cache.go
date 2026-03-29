package domain

import (
	"errors"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type CacheValue struct {
	Value    []byte
	ExpireAt time.Time
}
