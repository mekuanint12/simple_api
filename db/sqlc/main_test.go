package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root@:26257/users?application_name=%24+cockroach+sql&connect_timeout=15&sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to the database:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}

// import (
// 	"database/sql"
// 	"log"
// 	"os"
// 	"testing"

// 	_ "github.com/lib/pq"
// 	"github.com/techschool/simplebank/util"
// )

// var testQueries *Queries
// var testDB *sql.DB

// func TestMain(m *testing.M) {
// 	config, err := util.LoadConfig("../..")
// 	if err != nil {
// 		log.Fatal("cannot load config:", err)
// 	}

// 	testDB, err = sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	testQueries = New(testDB)

// 	os.Exit(m.Run())
// }
