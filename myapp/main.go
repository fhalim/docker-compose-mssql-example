package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	log.Println("Starting server on port 8080")
	hostname := os.Getenv("SQLSERVER_HOSTNAME")
	connString := fmt.Sprintf("sqlserver://sa:Pa55w0rd12345@%v?database=master", hostname)
	log.Println("Connection string: ", connString)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pool, err := sql.Open("sqlserver", connString)
		if err != nil {
			log.Fatal(err)
		}
		defer pool.Close()

		var t string
		err = pool.QueryRowContext(ctx, "SELECT table_name FROM information_schema.tables").Scan(&t)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "Hello, %q", t)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
