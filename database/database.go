package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Connection *pgxpool.Pool

// Connect function
func Initialize() (err error) {
	Connection, err = pgxpool.New(context.Background(), "postgres://username:password@localhost:5432/dbName")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	oniChan := make(chan bool, 1)
	go func(ch chan bool) {
		Connection.Ping(context.Background())
		ch <- true
	}(oniChan)

	select {
	case <-ctx.Done():
		fmt.Fprintln(os.Stderr, "Database Connection Timeout")
		os.Exit(1)
	case <-oniChan:
		log.Println("Database Connection Established!")
	}

	return nil
}
