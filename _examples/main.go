package main

import (
	"github.com/belldata/bdx"
	"github.com/belldata/bdx/interfaces"
)

func main() {
	r := bdx.New()
	r.GET("/health", func(c interfaces.Context) {
		c.JSON(200, bdx.B{"data": "test"})
	})
	r.Run()
}
