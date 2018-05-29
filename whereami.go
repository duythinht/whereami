package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	fmt.Println("Bootstrap Where Am I server...")

	var count int64

	client, err := statsd.NewClient("108.61.126.225:8125", "")

	if err != nil {
		log.Fatal(err)
	}

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Use(func(c martini.Context, req *http.Request, log *log.Logger, res http.ResponseWriter) {
		begin := time.Now().Nanosecond() / 1000

		c.Next()

		end := time.Now().Nanosecond() / 1000
		// Wrap http.ResponseWrite to martiti.ResponseWrite
		rw := res.(martini.ResponseWriter)

		telemetryKey := fmt.Sprintf("whereami.response.%s.%s.%v", req.Method, req.URL.Path, rw.Status())
		client.Timing(telemetryKey, int64(end-begin), 1.0)
	})

	m.Get("/", func(r render.Render) {
		hostname, _ := os.Hostname()
		count++
		r.JSON(200, map[string]interface{}{
			"ok": true,
			"data": map[string]interface{}{
				"Hostname": hostname,
				"Count":    count,
				"Version":  4,
			},
		})
	})

	m.Run()
}
