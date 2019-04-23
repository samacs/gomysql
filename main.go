package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		errc = make(chan error)

		hostname = "mysql"
		port     = 3306
		username = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		database = os.Getenv("MYSQL_DATABASE")

		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?autocommit=true&parseTime=true",
			username, password, hostname, port, database,
		)
	)
	log.Printf("DSN: %s", dsn)
	time.Sleep(5 * time.Second)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error getting database connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging the database: %v", err)
	}
	log.Printf("Connected successfully to %s", dsn)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		errc <- fmt.Errorf("%v", <-c)
	}()

	if err := <-errc; err != nil {
		log.Fatal(err)
	}
}
