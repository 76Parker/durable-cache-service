package api

import (
	 "context"

	 v1 "github.com/76Parker/durable-cache-service/gen/api/cache/v1"
	 "github.com/76Parker/durable-cache-service/internal/service"
	 "google.golang.org/protobuf/types/known/timestamppb"
)

type CacheServerImpl struct {
	 cache service.Cache
}

func (c *CacheServerImpl) Set(ctx context.Context, req *v1.SetRequest) (*v1.SetResponse, error) {
	 key := req.Key
	 value := req.Data
	 ttl := req.Ttl
	 expireAt := c.cache.Set(key, value, ttl.AsDuration())

	 return &v1.SetResponse{
		  ExpireAt: timestamppb.New(expireAt),
	 }, nil
}
func (c *CacheServerImpl) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	 return nil, nil
}

func NewCacheServer() v1.CacheServer {
	 return nil
}
