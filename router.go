package main

import (
	"github.com/qiangxue/fasthttp-routing"
)

func setRoutes(router *routing.Router) {

	router.Get("/twitter/<hash>", func(context *routing.Context) error {
		return handleTwitter(context)
	})

}
