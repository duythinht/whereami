package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	fmt.Println("Bootstrap Where Am I server...")

	var count int64
	client, err := statsd.NewClient("108.61.126.255:8125", "whereami")

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	client.Inc("stat1", 42, 1.0)

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Use(func(c martini.Context, req *http.Request, log *log.Logger) {
		c.Next()
		client.Timing("response.GET./.200", count, 1.0)
		log.Println("After request	")
	})

	m.Get("/", func(r render.Render) {
		hostname, _ := os.Hostname()
		count++
		r.JSON(200, map[string]interface{}{"ok": true, "data": map[string]interface{}{"hostname": hostname, "count": count, "version": 1}})
	})

	m.Run()
}
