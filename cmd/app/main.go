package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/rogue0026/door-locker/internal/application"
	"github.com/rogue0026/door-locker/internal/config"
)

func main() {
	appCfg := config.MustLoad()
	DSN := fmt.Sprintf("postgres://%s:%s@%s/%s", appCfg.DBUser, appCfg.DBUserPassword, appCfg.DBHost, appCfg.DatabaseName)
	pool, err := application.NewConnectionPool(context.Background(), DSN)
	if err != nil {
		panic(err)
	}
	app, err := application.New(appCfg, pool)
	if err != nil {
		panic(err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := app.Run()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println("server gracefully stopped")
				return
			} else {
				fmt.Printf("error occured while attempting gracefully stop: %s", err.Error())
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		<-stop
		fmt.Println("a stop signal has been received, trying to stop application gracefully")
		err := app.HTTPServer.Shutdown(context.Background())
		if err != nil {
			fmt.Printf("error occured while shutdown http server: %s", err.Error())
		}
		fmt.Println("closing database connection")
		app.CloseDatabaseConnection()
		fmt.Println("database connection closed")
	}()
	wg.Wait()
	fmt.Println("app terminated")
}
