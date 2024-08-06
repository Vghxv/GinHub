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
	dbtype := setting.DatabaseSetting.Type
	db := setting.DatabaseSetting.User
	password := setting.DatabaseSetting.Password
	host := setting.DatabaseSetting.Host
	port := setting.DatabaseSetting.Port
	name := setting.DatabaseSetting.Name
	databaseUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", dbtype, db, password, host, port, name)
	var err error
	DB, err = pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}
