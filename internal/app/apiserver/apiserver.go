package apiserver

import (
	"context"
	"github.com/MeguMan/mx_test/internal/app/store/postgres_store"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

func Start(dbConfig *Config) error{
	conn, err := newDB(dbConfig.DatabaseURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	s := postgres_store.New(conn)
	server := NewServer(*s)
	return http.ListenAndServe(":8080", server)
}

func newDB(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}