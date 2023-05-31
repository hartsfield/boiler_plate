package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	logFile := setupLogging()
	defer logFile.Close()

	srv := serverFromConf()
	ctx, cancelCtx := context.WithCancel(context.Background())

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
		cancelCtx()
	}()

	fmt.Println("Server started @ http://localhost" + srv.Addr)
	log.Println("Server started @ " + srv.Addr)
	<-ctx.Done()
}
