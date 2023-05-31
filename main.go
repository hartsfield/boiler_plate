package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// ckey/ctxkey is used as the key for the HTML context and is how we retrieve
// token information and pass it around to handlers
type ckey int

const (
	ctxkey ckey = iota
)

var (
	servicePort = os.Getenv("servicePort")
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	logFile := setupLogging()
	defer logFile.Close()

	mux.HandleFunc("/", home)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	srv := &http.Server{
		Addr:              servicePort,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	ctx, cancelCtx := context.WithCancel(context.Background())

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
		cancelCtx()
	}()

	log.Println("Server started @ " + srv.Addr)
	<-ctx.Done()
}
