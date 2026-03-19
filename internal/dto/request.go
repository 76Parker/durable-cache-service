package dto

import "time"

type SetRequest struct {
	Key   string
	Value []byte
	TTL   time.Duration
}
