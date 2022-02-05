package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Trey2k/OpenStreaming/webApp/app/common"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Create any tables that do not exist.
func init() {
	conn, err := connectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	makeViewers := false
	makeOverlays := false
	makeOverlayModules := false
	makeUsers := false

	_, err = conn.Query(context.Background(), "SELECT * FROM public.viewers")
	if err != nil {
		makeViewers = true
	}

	_, err = conn.Query(context.Background(), "SELECT * FROM public.overlays")
	if err != nil {
		makeOverlays = true
	}

	_, err = conn.Query(context.Background(), "SELECT * FROM public.\"overlayModules\"")
	if err != nil {
		makeOverlayModules = true
	}

	_, err = conn.Query(context.Background(), "SELECT * FROM public.users")
	if err != nil {
		makeUsers = true
	}

	if makeViewers {
		err = createViewersTable()
		if err != nil {
			common.Loggers.Error.Fatalf("Error while creating viewers table:\n%s\n", err)
		}
	}

	if makeOverlays {
		err = createOverlaysTable()
		if err != nil {
			common.Loggers.Error.Fatalf("Error while creating overlays table:\n%s\n", err)
		}
	}

	if makeOverlayModules {
		err = createOverlayModulesTable()
		if err != nil {
			common.Loggers.Error.Fatalf("Error while creating overlayModules table:\n%s\n", err)
		}
	}

	if makeUsers {
		err = createUsersTable()
		if err != nil {
			common.Loggers.Error.Fatalf("Error while creating users table:\n%s\n", err)
		}
	}

}

func connectDB() (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_DB")))
	if err != nil {
		return nil, err
	}

	return pool, nil

}

func createUsersTable() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

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
	defer conn.Close()

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
	defer conn.Close()

	sql, err := ioutil.ReadFile("/root/resources/sql/overlays.sql")
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), string(sql))
	return err
}

func createOverlayModulesTable() error {
	conn, err := connectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql, err := ioutil.ReadFile("/root/resources/sql/overlayModules.sql")
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), string(sql))
	return err
}
