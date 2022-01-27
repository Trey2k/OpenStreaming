package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Trey2k/OpenStreaming/app/common"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

// Create any tables that do not exist.
func init() {
	conn, err := connectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	common.Loggers.Info.Printf("Viewers table being reconstructed\n")
	_, err = conn.Query(context.Background(), "DROP TABLE IF EXISTS public.viewers")
	if err != nil {
		common.Loggers.Error.Fatalf("Error while querying DB:\n%s\n", err)
	}

	err = createViewersTable()
	if err != nil {
		common.Loggers.Error.Fatalf("Error while creating viewers table:\n%s\n", err)
	}

	common.Loggers.Info.Printf("Users table being reconstructed\n")
	_, err = conn.Query(context.Background(), "DROP TABLE IF EXISTS public.users")
	if err != nil {
		common.Loggers.Error.Fatalf("Error while querying DB:\n%s\n", err)
	}

	err = createUsersTable()
	if err != nil {
		common.Loggers.Error.Fatalf("Error while creating users table:\n%s\n", err)
	}

	common.Loggers.Info.Printf("Overlays table being reconstructed\n")
	_, err = conn.Query(context.Background(), "DROP TABLE IF EXISTS public.overlays")
	if err != nil {
		common.Loggers.Error.Fatalf("Error while querying DB:\n%s\n", err)
	}

	err = createOverlaysTable()
	if err != nil {
		common.Loggers.Error.Fatalf("Error while creating overlays table:\n%s\n", err)
	}

}

func connectDB() (*pgxpool.Pool, error) {
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

func createUsersTable() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}

	sql, err := ioutil.ReadFile("/root/resources/sql/users.sql")
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), string(sql))
	return err
}

func createViewersTable() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}

	sql, err := ioutil.ReadFile("/root/resources/sql/viewers.sql")
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), string(sql))
	return err
}

func createOverlaysTable() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}

	sql, err := ioutil.ReadFile("/root/resources/sql/overlays.sql")
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), string(sql))
	return err
}
