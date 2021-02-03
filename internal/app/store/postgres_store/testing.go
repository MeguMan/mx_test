package postgres_store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strings"
	"testing"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
}

func TestDB(t *testing.T) (*pgx.Conn, func(...string)) {
	t.Helper()
	dbConfig := &Config{}
	configFile, err := os.Open("../../../../configs/config.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(dbConfig); err != nil {
		fmt.Println(err)
	}

	conn, err := pgx.Connect(context.Background(), dbConfig.DatabaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	return conn, func(tables ...string) {
		if len(tables) > 0 {
			conn.Exec(context.Background(), fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		conn.Close(context.Background())
	}
}
