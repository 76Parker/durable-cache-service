package service

import (
	 "strings"
	 "sync"
	 "time"

	 "github.com/76Parker/durable-cache-service/internal/domain"
)

type Cache interface {
	 Set(key string, value []byte, ttl time.Duration) (expireAt time.Time)
	 Get(key string) (domain.CacheValue, error)
}

type cacheImpl struct {
	 m          *sync.Map
	 expireHeap *ExpireHeap
	 mu         *sync.Mutex
	 wake       chan struct{}
}

func NewCache(baseCapacity int) Cache {
	 c := &cacheImpl{
		  expireHeap: NewExpireHeap(baseCapacity),
		  mu:         new(sync.Mutex),
		  wake:       make(chan struct{}, 1),
		  m:          new(sync.Map),
	 }
	 go c.deleteScheduler()
	 return c
}
func (c *cacheImpl) Set(key string, value []byte, ttl time.Duration) (expireAt time.Time) {
	 expiring := time.Now().Add(ttl)
	 k := strings.TrimSpace(key)
	 v := domain.CacheValue{
		  Value:    value,
		  ExpireAt: expiring,
	 }
	 c.m.Store(k, v)
	 c.mu.Lock()
	 needWake := len(c.expireHeap.heap) == 0 || expiring.Before(time.Unix(0, c.expireHeap.heap[0].expireAt))
	 c.expireHeap.Push(k, expiring.UnixNano())
	 c.mu.Unlock()
	 if needWake {
		  select {
		  case c.wake <- struct{}{}:
		  default:
		  }
	 }
	 return expiring
}
func (c *cacheImpl) Get(key string) (value domain.CacheValue, err error) {
	 k := strings.TrimSpace(key)
	 v, ok := c.m.Load(k)
	 if !ok {
		  return domain.CacheValue{}, domain.ErrKeyNotFound
	 }
	 return v.(domain.CacheValue), nil
}

func (c *cacheImpl) deleteScheduler() {
	 var t *time.Timer
	 for {
		  c.mu.Lock()
		  if len(c.expireHeap.heap) == 0 {
				c.mu.Unlock()
				<-c.wake
				continue
		  }
		  nearest := c.expireHeap.heap[0]
		  wait := time.Until(time.Unix(0, nearest.expireAt))
		  if t == nil {
				t = time.NewTimer(wait)
		  } else {
				if !t.Stop() {
					 select {
					 case <-t.C:
					 default:
					 }
				}
				t.Reset(wait)
		  }
		  c.mu.Unlock()

		  select {
		  case <-t.C:
				now := time.Now()
				c.mu.Lock()
				for len(c.expireHeap.heap) > 0 && !time.Unix(0, c.expireHeap.heap[0].expireAt).After(now) {
					 expKey, ok := c.expireHeap.Pop()
					 if ok {
						  c.m.Delete(expKey)
					 }
				}
				c.mu.Unlock()
		  case <-c.wake:
		  }
	 }
}
