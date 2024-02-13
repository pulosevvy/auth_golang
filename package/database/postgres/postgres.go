package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type PGConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SslMode  string
}

func New(info *PGConfig) (*pgx.Conn, error) {
	const fn = "postgres.New"

	conn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		info.Username, info.Password, info.Host, info.Port, info.DBName, info.SslMode,
	)

	db, err := pgx.Connect(context.Background(), conn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	return db, err
}
