package main

import (
	"fmt"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	fmt.Println("Bootstrap Where Am I server...")

	var count int64

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		hostname, _ := os.Hostname()
		count++
		r.JSON(200, map[string]interface{}{
			"ok": true,
			"data": map[string]interface{}{
				"Hostname": hostname,
				"Count":    count,
				"Version":  5,
			},
		})
	})

	m.Run()
}
