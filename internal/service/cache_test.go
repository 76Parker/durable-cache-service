package service

import (
	"testing"
	"time"

	"github.com/76Parker/durable-cache-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestCacheImpl_Set(t *testing.T) {
	cache := NewCache(1000)
	key := "test_1"
	value := []byte("test_1")
	ttl := 5 * time.Second
	valueLen := len(value)
	_, err := cache.Set(key, value, ttl)
	assert.Nil(t, err)
	k, err := cache.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, valueLen, len(k.Value))
	time.Sleep(6 * time.Second)
	k, err = cache.Get(key)
	assert.IsType(t, domain.ErrKeyNotFound, err)
}

func TestCacheImpl_Get(t *testing.T) {
	cache := NewCache(1000)
	key := "test_1"
	value := []byte("test_1")
	ttl := 5 * time.Second
	_, err := cache.Set(key, value, ttl)
	assert.Nil(t, err)
	k, err := cache.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, []byte("test_1"), k.Value)
}
