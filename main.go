package main

import (
	"runtime"
	"strconv"

	"github.com/singcl/go-vue/routes"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pbnjay/memory"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	// 静态资源设置
	r.Static("/css", "/public/css")
	r.Static("/js", "/public/js")

	r.GET("/api/specs", func(c *gin.Context) {
		c.JSON(200, []string{
			runtime.GOOS,
			strconv.Itoa(runtime.NumCPU()),
			strconv.FormatUint(memory.TotalMemory()/(1024*1024), 10),
		})
	})

	r.POST("/api/persist", routes.Persist)
	r.POST("/api/descriptive/mean", routes.Mean)
	r.POST("/api/descriptive/standard-deviation", routes.StdDev)
	r.POST("/api/distribution", routes.NormalCDF)

	// r.Run(":3000")
	r.Run()
}
