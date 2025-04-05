package db

import (
	"context"

	"product_service/utils"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Connect(c *utils.DatabaseConfig) (err error) {
	var connection_str string = "postgresql://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.DBName
	conn, err = pgx.Connect(context.Background(), connection_str)
	return err
}
func Close() error {
	if conn != nil {
		return conn.Close(context.Background())
	}
	return nil
}
