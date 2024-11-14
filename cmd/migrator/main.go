package main

import (
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	migrationPath string = "file://migrations/"
	connConfig    string = "pgx5://user:password@localhost:5432/door_locks"
	direction     string
)

func main() {
	flag.StringVar(&direction, "direction", "", "migration direction: up or down")
	flag.Parse()

	inst, err := migrate.New(migrationPath, connConfig)
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
