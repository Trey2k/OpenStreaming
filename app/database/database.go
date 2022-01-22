package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// Create any tables that do not exist.
func init() {
	conn, err := connectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	_, err = conn.Query(context.Background(), "select * from viewers")
	if err != nil {
		log.Println("Viewers table not found - creating now")
		err = createViewersTable()
		if err != nil {
			panic(err)
		}
	}

}

func connectDB() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_DB")))

}
