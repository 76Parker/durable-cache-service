package service

import (
	 "sync"
)

type Cache struct {
	 m sync.Map
}

func NewCache() *Cache {
	 return &Cache{
		  sync.Map{},
	 }
}
