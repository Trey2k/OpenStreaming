package database

import (
	"context"
	"io/ioutil"
)

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
