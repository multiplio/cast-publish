package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/qiangxue/fasthttp-routing"
)

func (sc *serverContext) handleTwitter(c *routing.Context) error {
	// get params
	userID := c.Param("user")
	if userID == "" {
		log.Println("user missing")
		return routing.NewHTTPError(400, "No user.")
	}
	postID := c.Param("post")
	if postID == "" {
		log.Println("post missing")
		return routing.NewHTTPError(400, "No post.")
	}

	log.Println("got userID :", userID, "and postID :", postID)

	// get user secret and token
	var user struct {
		twitter struct {
			token  string
			secret string
		}
	}
	filter := bson.M{"_id": userID}
	err := sc.users.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)

	// auth with twitter

	return nil
}
