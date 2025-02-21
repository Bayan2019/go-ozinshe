package configuration

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Bayan2019/go-ozinshe/repositories"
	_ "github.com/mattn/go-sqlite3"
)

// var MapCfg *MapConfiguration
var ApiCfg *ApiConfiguration

// type MapConfiguration struct{}

type ApiConfiguration struct {
	Conn      *sql.DB
	DB        *repositories.Queries
	DirImages string
	DirVideos string
}

func Connect2DB(dbPath, platform string) error {
	// https://github.com/libsql/libsql-client-go/#open-a-connection-to-sqld
	// libsql://[your-database].turso.io?authToken=[your-auth-token]
	// dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("inside Connect2DB")
	fmt.Println(dbPath)
	if dbPath == "" {
		return errors.New("No DataBase Path")
	} else {
		var db *sql.DB

		if platform == "dev" {
			db1, err := sql.Open("sqlite3", dbPath)
			if err != nil {
				return err
			}
			db = db1
		} else {
			db1, err := sql.Open("libsql", dbPath)
			if err != nil {
				return err
			}
			db = db1
		}

		dbQueries := repositories.New(db)
		ApiCfg = &ApiConfiguration{
			Conn: db,
			DB:   dbQueries,
		}
		log.Println("Connected to database!")
		return nil
	}
}
