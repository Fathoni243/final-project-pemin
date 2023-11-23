package handler

import (
	"final-project-pemin/database"
	"final-project-pemin/middleware"
	"final-project-pemin/src/service"
	"final-project-pemin/util"
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
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token"},
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

	r.httpServer.POST("/seeder", func(c *gin.Context) {
		database.SeederRefresh()

		util.SuccessResponse(c, http.StatusOK, "seeder success", "data", nil)
	})

	prodi := r.httpServer.Group("/prodi")
	{
		prodi.GET("", r.FindAllProdiHandler)
	}

	auth := r.httpServer.Group("/auth")
	{
		auth.POST("/register", r.RegisterMahasiswa)
		auth.POST("/login", r.LoginMahasiswa)
	}

	mahasiswa := r.httpServer.Group("/mahasiswa")
	{
		mahasiswa.GET("/:nim", r.FindMahasiswaByNIM)
		mahasiswa.GET("", r.FindAllMahasiswa)
		mahasiswa.GET("/profile", middleware.ValidateToken(), r.GetProfileMahasiswa)
		mahasiswa.POST("/matakuliah/:mkId", middleware.ValidateToken(), r.SaveMatkulMahasiswa)
		mahasiswa.PUT("/matakuliah/:mkId", middleware.ValidateToken(), r.DeleteMatkulMahasiswa)
	}

	mataKuliah := r.httpServer.Group("/matakuliah")
	{
		mataKuliah.GET("", r.FindAllMataKuliah)
	}
}

func (r *rest) Run() {
	if err := r.httpServer.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
