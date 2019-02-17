package main

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/qiangxue/fasthttp-routing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

	// get post

	// get user secret and token
	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("could not convert", userID, "into ObjectID")
		return routing.NewHTTPError(400, "Invalid user.")
	}

	var user struct {
		Twitter struct {
			Token  string
			Secret string
		}
	}
	filter := bson.M{"_id": userOID}
	err = sc.users.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println("could not find user", userID)
		return routing.NewHTTPError(400, "Invalid user.")
	}

	// auth with twitter
	client, err = authTwitter(&environment.Twitter.Key, &environment.Twitter.Secret, &user.Twitter.Token, &user.Twitter.Secret)
	if err != nil {
		log.Println("could not authenticate", userID)
		return routing.NewHTTPError(400, "Invalid user.")
	}

	return nil
}

func authTwitter(consumerKey, consumerSecret, accessToken, accessSecret *string) (*twitter.Client, error) {
	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return client, nil
}
