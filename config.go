package main

import (
	"fmt"
	"os"
)

var (
	pgHost, pgDB, pgUser, pgPassword string
	svcPort, svcUser, svcPassword    string
)

func init() {

	pgHost = os.Getenv("PG_HOST")
	pgDB = os.Getenv("PG_DB")
	pgUser = os.Getenv("PG_USER")
	pgPassword = os.Getenv("PG_PASSWORD")

}

// ConnectSting construct database connection string
func ConnectSting() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s)", pgUser, pgPassword, pgHost, pgDB)
}
