package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jasurxaydarov/marifat_ac_backend/api/docs"
	"github.com/jasurxaydarov/marifat_ac_backend/api/handlers"
	"github.com/jasurxaydarov/marifat_ac_backend/redis"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Api(storage storage.StorageI, cache redis.RedisRepoI) {

	router := gin.Default()

	h := handlers.NewHandlers(storage, cache)

	router.GET("/ping")

	router.POST("/user", h.UserCreate)
	router.POST("/user-check", h.CheckUser)
	router.POST("/singup", h.SignUp)
	router.POST("/login", h.LogIn)
	router.GET("/user/:id", h.GetUserById)
	router.POST("/users", h.GetUsers)
	router.POST("/for_req", h.CreateForReq)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	router.Run(":8080")
}
