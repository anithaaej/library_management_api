package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anithaa19/bms/routers"
	// "github.com/go-sql-driver/mysql"
)

// var db *sql.DB

// func dbConfig() {
// 	cfg := mysql.Config{
// 		User:   os.Getenv("root"),
// 		Passwd: os.Getenv("Root@123"),
// 		Net:    "tcp",
// 		Addr:   "127.0.0.1:3306",
// 		DBName: "bms",
// 	}
// 	fmt.Println(cfg.FormatDSN())
// 	// Get a database handle.
// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }

func main() {
	fmt.Println("Welcome to Book management system")
	// dbConfig()

	r := routers.Routers()

	fmt.Println("Server is getting Started..")

	log.Fatal(http.ListenAndServe(":3000", r))

	fmt.Println("Listening at Port: 3000")
}
