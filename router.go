package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/qiangxue/fasthttp-routing"
)

type serverContext struct {
	users *mongo.Collection
}

func setRoutes(router *routing.Router, sc *serverContext) {

	router.Get("/twitter/<user>/<post>", func(context *routing.Context) error {
		return sc.handleTwitter(context)
	})

}
