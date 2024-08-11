package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Vghxv/GinHub/pkg/setting"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	config := setting.DatabaseSetting
	databaseURL := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		config.Type,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)	

	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}

	DB, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	log.Println("Database connected")
}

func Close() {
	DB.Close()
	log.Println("Database closed")
}
