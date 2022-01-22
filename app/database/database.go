package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

// Create any tables that do not exist.
func init() {
	conn, err := ConnectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Query(context.Background(), "SELECT * FROM viewers")
	if err != nil && err != pgx.ErrNoRows {
		log.Println("Viewers table not found - creating now")
		err = createViewersTable()
		if err != nil {
			panic(err)
		}
	}

	_, err = conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil && err != pgx.ErrNoRows {
		log.Println("Users table not found - creating now")
		err = createUsersTable()
		if err != nil {
			panic(err)
		}
	}

}

func ConnectDB() (*pgxpool.Pool, error) {
	if Pool == nil {
		var err error
		Pool, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("POSTGRES_DB")))
		if err != nil {
			return nil, err
		}
	}
	return Pool, nil

}
