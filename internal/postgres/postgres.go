package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitPoolConnection() (*pgxpool.Pool, error) {

	//user := getEnvs("DATABASE_USER")
	//	password := getEnvs("DATABASE_PASSWORD")
	//dbname := getEnvs("DATABASE_DBNAME")
	//port := getEnvs("DATABASE_PORT")
	//host := getEnvs("DATABASE_HOST")
	//connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://zucrqvwk:b0EDhSDeZcQS9KAiXfHLR-G8tKt_vL9r@batyr.db.elephantsql.com/zucrqvwk")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		//	log.Println(user, password, dbname, port, host)
		os.Exit(1)
	}

	return dbpool, nil

}

/*func getEnvs(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}*/
