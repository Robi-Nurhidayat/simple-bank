package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

//func TestMain(m *testing.M) {
//	conn, err := sql.Open(dbDriver, dbSource)
//
//	if err != nil {
//		log.Fatal("cannot connect to db", err)
//	}
//
//	testQueries = New(conn)
//
//	os.Exit(m.Run())
//}

//func TestConn(t *testing.T) {
//	conn, err := sql.Open(dbDriver, dbSource)
//
//	if err != nil {
//		log.Fatal("cannot connect to db", err)
//	}
//
//	testQueries = New(conn)
//
//	fmt.Println("suksess")
//	//os.Exit(m.Run())
//}
