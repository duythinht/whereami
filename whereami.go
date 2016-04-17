package main

import (
	"fmt"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var count int32

func main() {
	fmt.Println("Bootstrap Where Am I server...")

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		hostname, _ := os.Hostname()
		count++
		r.JSON(200, map[string]interface{}{"ok": true, "hostname": hostname, "count": count})
	})

	m.Run()
}
