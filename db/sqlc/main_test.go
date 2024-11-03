package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/adiubaidah/rfid-syafiiyah/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	logrus := config.NewLogger()
	config, err := config.LoadEnv("../..")
	log.SetOutput(logrus.Writer())
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSourceTest)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(connPool)

	os.Exit(m.Run())
}
