package database

import (
	"context"
	"io/ioutil"
)

func createUsersTable() error {
	conn, err := ConnectDB()
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
