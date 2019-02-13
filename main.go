package main

import (
	"context"
	"log"
	"time"

	env "github.com/Netflix/go-env"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type environmentDesc struct {
	Alias   string `env:"PROGRAM_ALIAS"`
	Address string `env:"ADDRESS"`

	Twitter struct {
		Key    string `env:"TWITTER_CONSUMER_KEY"`
		Secret string `env:"TWITTER_CONSUMER_SECRET"`
	}

	Users struct {
		Name     string `env:"DATABASE_NAME"`
		Address  string `env:"DATABASE_ADDRESS"`
		Options  string `env:"DATABASE_OPTIONS"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
	}
}

var environment environmentDesc

func main() {
	log.SetFlags(0)

	// get environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	log.SetPrefix(environment.Alias + ":")

	// server context
	connectionString := `mongodb://` + environment.Users.Name + `:` + environment.Users.Password + `@` + environment.Users.Address + `/` + environment.Users.Name
	if environment.Users.Options != "" {
		connectionString += `?` + environment.Users.Options
	}

	client, err := mongo.NewClient(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	context := serverContext{
		users: client.Database(environment.Users.Name).Collection("users"),
	}

	// routes
	router := routing.New()
	setRoutes(router, &context)

	// start server
	log.Println("Serving at ", environment.Address)
	log.Fatal(fasthttp.ListenAndServe(environment.Address, router.HandleRequest))
}
