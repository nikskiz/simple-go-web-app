package main

import (
	"log"
	"fmt"
	"net/http"
	"database/sql"
	 "os"
	// GITHUB
	_ "github.com/lib/pq"

)

func containsEmpty(ss ...string) bool {
    for _, s := range ss {
        if s == "" {
            return true
        }
    }
    return false
}

func main() {
	server := os.Getenv("AWS_RDS_HOSTNAME")
	user := os.Getenv("AWS_RDS_USERNAME") //TODO Use unicreds or credstash to retreieve creds
	password := os.Getenv("AWS_RDS_PASSWORD") //TODO Use unicreds or credstash to retreieve creds
	database := os.Getenv("DATABASE_NAME")

	if containsEmpty(server,user,password,database) {
   log.Fatal("ERROR - Ensure ENV VARS exist AWS_RDS_HOSTNAME AWS_RDS_USERNAME AWS_RDS_PASSWORD DATABASE_NAME")
 	}
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user,password,server,database)

		// Connect to the DB, panic if failed
		// TODO Use os env vars to store creds and hostname
	db, err := sql.Open("postgres", connString)
	if err != nil {
			fmt.Println(`Could not connect to db`)
			panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT payment_id, product_name, amount FROM payments WHERE payment_id=1`)
	if err != nil {
	    panic(err)
	}
	var (
		payment_id int
		product_name string
		amount float64
	)
	for rows.Next() {
	    rows.Scan(&payment_id, &product_name,&amount)
	    fmt.Println(payment_id, product_name, amount)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey!, you've requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "You have purchased a %s at $%f\n", product_name, amount)
	})

	defaultPort := "8080"
	log.Printf("Listening on default port: %s", defaultPort)
	log.Fatal(http.ListenAndServe(":" + defaultPort, nil))
}
