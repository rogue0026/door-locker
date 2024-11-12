package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/rogue0026/door-locker/internal/application"
	"github.com/rogue0026/door-locker/internal/config"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var (
	cfgPath string
)

func main() {
	flag.StringVar(&cfgPath, "cfg", "", "path to application config")
	flag.Parse()

	appCfg := config.MustLoad(cfgPath)
	appStorage, err := postgres.New(context.Background(), appCfg.DSN)
	if err != nil {
		panic(err)
	}
	app := application.New(appCfg, &appStorage)
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
		app.AppStorage.Close()
		fmt.Println("database connection closed")
	}()
	wg.Wait()
	fmt.Println("app terminated")
}
