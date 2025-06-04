package database

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

var connStr string = "postgres://root:root@localhost:5432/booker"

func GetConnection() (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, errors.New("cannot connect to postgres")
	}

	return conn, nil
}
