package handlers

import (
	"github.com/jasurxaydarov/marifat_ac_backend/redis"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
)

type Handlers struct {
	storage storage.StorageI
	cache   redis.RedisRepoI
}

func NewHandlers(storage storage.StorageI, cache redis.RedisRepoI) Handlers {

	return Handlers{storage: storage, cache: cache}

}
