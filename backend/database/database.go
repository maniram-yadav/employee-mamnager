package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

func DbConn() *sql.DB {
	log.Println("inside DbConn()")
	if err := godotenv.Load(); err != nil {
		log.Println(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	dbHost, errHost := os.LookupEnv("DB_HOST")
	dbUser, errUser := os.LookupEnv("DB_USER")
	dbPasword, errPass := os.LookupEnv("DB_PASS")

	if !errHost {
		log.Println(errors.New("DB_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	if !errUser {
		log.Println(errors.New("DB_USER not set in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	if !errPass {
		log.Println(errors.New("DB_PASS not set in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", dbUser, dbPasword, dbHost)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		log.Println(err)
		log.Fatal(err)
	} else {
		log.Println("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE IF NOT EXISTS employees`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE employees`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE employee (
		    empId varchar(20) NOT NULL,
		    name varchar(20) NOT NULL,
		    email varchar(20) NOT NULL,
			contact varchar(20) NOT NULL,
		    PRIMARY KEY (empId)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}

func InitDatabase() {

}
