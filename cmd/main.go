package main

import (
	"context"
	"fmt"

	"github.com/AliAstanov/my_blogs/api"
	"github.com/AliAstanov/my_blogs/config"
	"github.com/AliAstanov/my_blogs/models"
	"github.com/AliAstanov/my_blogs/pkg/db"
	log "github.com/AliAstanov/my_blogs/pkg/logger"
	"github.com/AliAstanov/my_blogs/pkg/token"
	"github.com/AliAstanov/my_blogs/storage"
	"github.com/AliAstanov/my_blogs/storage/redis"
	"github.com/saidamir98/udevs_pkg/logger"
)

var ctx = context.Background()

func main() {
	cfg := config.Load()

	log := log.NewLogger(cfg.GeneralConfig)

	pgxPool, err := db.ConnDb(cfg.PgConfig)
	if err != nil {
		log.Error("error with connecting database:", logger.Error(err))
		return
	}
	log.Debug("soccessfully connected with db")

	fmt.Println(pgxPool)

	redisCli, err := db.ConnRedis(log, ctx, cfg.RedisConfig)
	if err != nil {
		log.Error("error with connecting redis:", logger.Error(err))
	}

	log.Debug("soccessfully connected with redis")

	cache := redis.NewRedisRepo(redisCli, log)

	r := redis.NewRedisRepo(redisCli, log)

	v, _ := r.Exists(ctx, "gov")

	fmt.Println(v)

	storage := storage.NewStorage(pgxPool, log)

	fmt.Println("str", storage)

	engine := api.Api(api.Options{
		Storage: storage,
		Log:     log,
		Cache:   cache,
	})

	fmt.Println(" ---------- ")

	tokens, err := token.GenerateJWT(models.Claim{UserId: "001",UserRole: "Ali"})
	if err != nil {
		fmt.Println(err)
	return
	}
	fmt.Println("token----->>",tokens)

	claim, err := token.ParsJwt(tokens)
	if err != nil {
		fmt.Println(err)
	return
	}
	fmt.Println(claim)

	log.Debug("server is running on ", logger.String("port:", cfg.GeneralConfig.HTTPPort))
	engine.Run(cfg.GeneralConfig.HTTPPort)


}
