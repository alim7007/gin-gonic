package main

import (
	"io"
	"net/http"
	"os"

	"github.com/alim7007/gin-gonic/controller"
	"github.com/alim7007/gin-gonic/middlewares"
	"github.com/alim7007/gin-gonic/service"
	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	// server := gin.Default()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	// server.Use(gin.Recovery(), gin.Logger())
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth()) //  gindump.Dump()

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/video", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video is Valid"})
			}

		})
	}
	veiwRoutes := server.Group("/view")
	{
		veiwRoutes.GET("/videos", videoController.ShowAll)
	}
	server.Run(":8001")
}
