package main

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/marifat_ac_backend/api"
	"github.com/jasurxaydarov/marifat_ac_backend/config"
	"github.com/jasurxaydarov/marifat_ac_backend/pgx/db"
	"github.com/jasurxaydarov/marifat_ac_backend/redis"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {

	fmt.Println("start")
	log:=logger.NewLogger("",logger.LevelDebug)

	
	cfg := config.Load()

	fmt.Println(cfg)

	pgxCoon, err := db.ConnectDB(cfg.PgConfig)

	if err != nil {
		fmt.Println("err on conn db", err)
		return
	}
	fmt.Println(pgxCoon)

	str := storage.NewStorage(pgxCoon)

	redisCli,err:=db.ConnRedis(context.Background(),cfg.RedisConfig)

	if err != nil {
		fmt.Println("err on conn redis", err)
		return
	}

	fmt.Println(redisCli)

	cache:=redis.NewRedisRepo(redisCli,log)

	api.Api(str,cache)
}
