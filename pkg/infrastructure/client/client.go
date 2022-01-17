package client

import (
	"StegoLSB/ent"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq" // So we can use dialect.Postgres
	"log"
	"os"
	"strconv"
)

func psqlInfo() string {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("invalid DB_PORT environment variable: %v", err)
	}

	return fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		dbPort,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"),
	)
}

// New returns a new instance of ent.Client.
func New() (*ent.Client, error) {
	var entOptions []ent.Option
	_ = append(entOptions, ent.Debug())

	return ent.Open(dialect.Postgres, psqlInfo(), entOptions...)
}
