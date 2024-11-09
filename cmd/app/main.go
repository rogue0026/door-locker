package main

import (
	"flag"
	"fmt"
	"github.com/rogue0026/door-locker/internal/application"
	"github.com/rogue0026/door-locker/internal/config"
)

var (
	cfgPath string
)

func main() {
	flag.StringVar(&cfgPath, "cfg", "", "path to application config")
	flag.Parse()

	appCfg := config.MustLoad(cfgPath)
	app := application.New(appCfg)
	address := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	err := app.Run(address)
	if err != nil {
		fmt.Println(err.Error())
	}
}
