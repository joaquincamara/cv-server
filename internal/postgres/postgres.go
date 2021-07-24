package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func InitPoolConnection() (*pgxpool.Pool, error) {

	user := GetEnvs("DATABASE_USER")
	password := GetEnvs("DATABASE_PASSWORD")
	dbname := GetEnvs("DATABASE_DBNAME")
	port := GetEnvs("DATABASE_PORT")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)
	if len(connectionString) == 45 {
		connectionString = "postgres://zucrqvwk:b0EDhSDeZcQS9KAiXfHLR-G8tKt_vL9r@batyr.db.elephantsql.com/zucrqvwk"
	}

	dbpool, err := pgxpool.Connect(context.Background(), connectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		log.Println(user, password, dbname, port)
		os.Exit(1)
	}

	return dbpool, nil

}

func GetEnvs(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
