package main

import (
	"CONTACTAPP/api"
	"CONTACTAPP/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)


func main() {
	connStr := "user=bahodir dbname=postgres password=7020 sslmode=disable"
	db,err := sql.Open("postgres",connStr)
	if err != nil{
		log.Println(err.Error())
		log.Fatal()
	}
	defer db.Close()
	service.InitializeDatabase(db)


	fmt.Println("Server is running on http://localhost:9999")

	// http.HandleFunc("/signup", api.SignUp)
	http.HandleFunc("/add", api.SignUp)
	http.HandleFunc("/getall", api.GetAllUsers)
	http.ListenAndServe("localhost:9999", nil)

}




