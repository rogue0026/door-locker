package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rogue0026/door-locker/internal/config"
)

var (
	appConfigPath string
	direction     string
)

func main() {
	flag.StringVar(&appConfigPath, "cfg", "", "path to application config file")
	flag.StringVar(&direction, "direction", "", "migration direction: up or down")
	flag.Parse()

	appConfig := config.MustLoad(appConfigPath)
	inst, err := migrate.New(appConfig.MigrationPath, appConfig.MigrationDSN)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch direction {
	case "up":
		err = inst.Up()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case "down":
		err = inst.Down()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	default:
		fmt.Println("invalid migration direction")
		return
	}
	fmt.Println("migrated successfully")
}
