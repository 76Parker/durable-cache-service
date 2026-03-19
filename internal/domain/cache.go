package domain

import "time"

type CacheValue struct {
	Value    []byte
	ExpireAt time.Time
}
