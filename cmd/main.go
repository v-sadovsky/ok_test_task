package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/v-sadovsky/ok_test_task/pkg/models"
	"github.com/v-sadovsky/ok_test_task/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	stores   interface {
		Get(string) (map[int]*models.Shop, error)
	}
	pool int
}

func main() {
	t := time.Now()
	curTime := fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
	dsn := fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	tm := flag.String("tm", curTime, "\"Current time\" for testing purpose")
	fName := flag.String("f", "orders.xml", "Name of the target XML data file")
	wp := flag.Int("wp", 6, "Workers pool size")
	flag.Parse()

	iLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	eLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	_, err := os.Stat(*fName)
	if err == nil {
		iLog.Printf("File %s already exists", *fName)
		return
	} else if !errors.Is(err, os.ErrNotExist) {
		eLog.Fatal(err)
	}

	db, err := openDB(dsn)
	if err != nil {
		eLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog:  iLog,
		errorLog: eLog,
		stores:   &mysql.StoresModel{DB: db},
	}

	data, err := app.requestData(*wp, *tm)
	if err == models.ErrNoRecord {
		app.infoLog.Println(err)
		return
	} else if err != nil {
		eLog.Fatal(err)
	}

	err = provideFile(*fName, data)
	if err != nil {
		eLog.Fatal(err)
	}

	iLog.Printf("Requested data have been written to the %s file.", *fName)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
