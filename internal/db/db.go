package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB() {
	config, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatal(err)
	}

	config.Host = "aws-1-ap-south-1.pooler.supabase.com"
	config.Port = 6543
	config.User = "postgres.plscuntonlxcybhmwfnl"
	config.Password = "MiT@#$271102"
	config.Database = "postgres"
	config.Config.RuntimeParams["sslmode"] = "require"

	Conn, err = pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	log.Println("✅ Connected to Supabase")
}
