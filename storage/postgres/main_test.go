package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/Muhammadjon226/toDo-service/config"
	"github.com/Muhammadjon226/toDo-service/pkg/db"
	"github.com/Muhammadjon226/toDo-service/pkg/logger"
)

var pgRepo *taskRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgRepo = NewTaskRepo(connDB)

	os.Exit(m.Run())
}
