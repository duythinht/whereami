package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	fmt.Println("Bootstrap Where Am I server...")

	var count int64
	client, _ := statsd.NewClient("108.60.126.255:8125", "test-client")

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Use(func(req http.Request) {
		client.Timing("response.GET./.200", count, 1.0)
	})

	m.Get("/", func(r render.Render) {
		hostname, _ := os.Hostname()
		count++
		r.JSON(200, map[string]interface{}{"ok": true, "data": map[string]interface{}{"hostname": hostname, "count": count, "version": 1}})
	})

	m.Run()
}
