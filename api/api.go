package api

import (
	_ "github.com/AliAstanov/my_blogs/api/docs"
	"github.com/AliAstanov/my_blogs/api/handlers"
	log "github.com/AliAstanov/my_blogs/pkg/logger"
	"github.com/AliAstanov/my_blogs/storage"
	"github.com/AliAstanov/my_blogs/storage/redis"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Options struct {
	Storage storage.StorageI
	Log     log.Log
	Cache   redis.RedisRepoI
}

func Api(opt Options) *gin.Engine {

	h := handlers.NewHandlers(handlers.Handlers{Storage: opt.Storage, Log: opt.Log, Cache: opt.Cache})

	engine := gin.Default()

	api := engine.Group("/api")

	api.GET("/ping", h.Ping)


	//owner
	own := api.Group("/own")
	{
		// own.POST("/log-in")
		// own.POST("/log-out")
		own.POST("/categorys", h.CreateCategories)
		// own.PUT("/update-category/:id")
		// own.DELETE("/delete-category/:id")
	}

	//viewer
	// vw := api.Group("/vw")
	{
		// vw.POST("/log-out")
		// vw.POST("/comment/:article_id")
	}

	//public
	pb := api.Group("/pb")
	{
		pb.POST("/check-user", h.CheckUser)
		// pb.POST("/sign-up")
		// pb.POST("/log-in")
		pb.GET("/categories", h.GetCategoriesList)
		pb.GET("/category/:id", h.GetCategory)
		// pb.GET("/articles")
		// pb.GET("/articles/:id")
	}

	url := ginSwagger.URL("swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return engine
}
