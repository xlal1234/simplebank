package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/xlal1234/simplebank/db/util"
)

// var testStore Store

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connPool, err := pgxpool.New(context.Background(), config.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }

	// testStore = NewStore(connPool)
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
