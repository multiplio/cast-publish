package main

import (
	"log"

	env "github.com/Netflix/go-env"
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

	// routes
	router := routing.New()
	setRoutes(router)

	// start server
	log.Println("Serving at ", environment.Address)
	log.Fatal(fasthttp.ListenAndServe(environment.Address, router.HandleRequest))
}
