package main

import (
	"embed"
	"home/leetcode"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var tmpl embed.FS

//go:embed static
var static embed.FS

func main() {
	r := gin.Default()
	t, _ := template.ParseFS(tmpl, "templates/*.html")
	r.SetHTMLTemplate(t)
	r.GET("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(static))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Leetcode": leetcode.LeetcodeInfo,
		})
	})
	api := r.Group("/api")
	api.GET("/leetcode", leetcode.GetInfo)
	r.Run()
}
