package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestInitPoolConnection(t *testing.T) {
	const (
		user     = "postgres"
		password = "Joaquinc1"
		dbName   = "curriculum"
		port     = "5432"
		host     = "localhost"
	)

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbName, port)
	_, ok := pgxpool.Connect(context.Background(), connectionString)

	if ok != nil {
		t.Fatal("Connection and pool Fail")
	}

}
