package handlers

import (
	log "github.com/AliAstanov/my_blogs/pkg/logger"
	"github.com/AliAstanov/my_blogs/storage"
	"github.com/AliAstanov/my_blogs/storage/redis"
	"github.com/gin-gonic/gin"
)

type handlers struct {
	storage storage.StorageI
	log     log.Log
	cache   redis.RedisRepoI
}

type Handlers struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func NewHandlers(h Handlers) handlers {
	return handlers{h.Storage, h.Log, h.Cache}
}

func (h *handlers) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message:": "Pong"})
}
