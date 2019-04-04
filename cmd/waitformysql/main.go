package main

import (
	"bytes"
	"flag"
	"fmt"
	"database/sql"
	"os"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	user := flag.String("username", "root", "database username")
	password := flag.String("password", "", "database password")
	host := flag.String("host", "127.0.0.1", "database host")
	port := flag.String("port", "3306", "database port")
	database := flag.String("database", "", "database name")
	timeout := flag.Duration("timeout", 25 * time.Second, "max timeout in seconds")

	flag.Parse()

	if *password == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	connectionUri := generateConnectionUri(*user, *password, *host, *port, *database)

	connection := getDb(connectionUri)

	c := make(chan bool, 1)

	go func() {
		for ! ping(connection) {}
		c <- true
	}()

	select {
		case <-c:
			connection.Close()
			os.Exit(0)
		case <-time.After(*timeout):
			connection.Close()
			os.Exit(1)
	}
}

func generateConnectionUri(username string, password string, host string, port string, database string) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%s:%s@", username, password))
	buffer.WriteString(fmt.Sprintf("tcp(%s:%s)", host, port))
	if database != "" {
		buffer.WriteString(fmt.Sprintf("/%s", database))
	}

	return buffer.String()
}

func getDb(connection string) *sql.DB {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err.Error())

	}

	return db;
}

func ping(connection *sql.DB) bool {

	err := connection.Ping()
	if err != nil {
		return false;
	}

	return true;
}