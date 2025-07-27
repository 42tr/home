package main

import (
	"embed"
	"home/leetcode"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var staticFiles embed.FS

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	assetFS, err := fs.Sub(staticFiles, "frontend/dist/assets")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/assets", http.FS(assetFS))

	// 处理 favicon
	r.GET("/star.png", func(c *gin.Context) {
		file, err := staticFiles.ReadFile("frontend/dist/star.png")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "image/x-icon", file)
	})
	// 处理首页
	r.GET("/", func(c *gin.Context) {
		file, err := staticFiles.ReadFile("frontend/dist/index.html")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	// SPA 路由处理 - 处理所有其他路由
	r.NoRoute(func(c *gin.Context) {
		// 如果是 API 请求，返回 404
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "API not found"})
			return
		}

		// 其他情况返回 index.html (SPA 支持)
		file, err := staticFiles.ReadFile("frontend/dist/index.html")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	api := r.Group("/api")
	api.GET("/leetcode", leetcode.GetInfo)
	r.Run(":9042")
}
