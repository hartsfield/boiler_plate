package main

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/redis/go-redis/v9"
)

type viewData struct{}

// ckey/ctxkey is used as the key for the HTML context and is how we retrieve
// token information and pass it around to handlers
type ckey int

const (
	ctxkey ckey = iota
)

var (
	servicePort = ":" + os.Getenv("servicePort")
	logFilePath = os.Getenv("logFilePath")

	redisIP = os.Getenv("redisIP")
	rdb     = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":6379",
		Password: "",
		DB:       0,
	})
	rdx = context.Background()

	templates = template.Must(template.New("main").ParseGlob("internal/*/*"))
	mux       = http.NewServeMux()
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	mux.HandleFunc("/", home)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
