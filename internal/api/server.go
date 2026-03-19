package api

import (
	"context"

	v1 "github.com/76Parker/durable-cache-service/gen/api/cache/v1"
)

type CacheServerImpl struct {
}

func (c *CacheServerImpl) Set(ctx context.Context, req *v1.SetRequest) (*v1.SetResponse, error) {
	return nil, nil
}
func (c *CacheServerImpl) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	return nil, nil
}

func NewCacheServer() v1.CacheServer {
	return
}
