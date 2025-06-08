package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB 

func InitDB(){
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	db,err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(localhost)/%s",DB_USER,DB_PASSWORD,DB_NAME))
	if err != nil {
		log.Panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Panic(err.Error())
	}
	Db = db 
}

func CloseDB() error {
	return Db.Close()
}

func Migrate(){
	if err := Db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	driver,_ := mysql.WithInstance(Db,&mysql.Config{})

	m,_ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange{
		log.Fatal(err)
	}
	
}