package handler

import (
	"final-project-pemin/src/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type rest struct {
	httpServer *gin.Engine
	service    *service.Service
}

func Init(service *service.Service) *rest {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	rest := &rest{
		httpServer: r,
		service:    service,
	}

	rest.RegisterMiddlewareAndRoutes()

	return rest
}

func (r *rest) RegisterMiddlewareAndRoutes() {
	r.httpServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test success",
		})
	})

	// r.httpServer.GET("/seeder", func(ctx *gin.Context) {
	// 	helper.SeederRefresh()

	// 	helper.ResponseSuccessJson(ctx, "seeder success", "")
	// })

	prodi := r.httpServer.Group("/prodi")
	{
		prodi.GET("", r.FindAllProdiHandler)
	}
}

func (r *rest) Run() {
	if err := r.httpServer.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
