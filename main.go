package main

import (
	"Go-Mux-Smm/artist"
	"Go-Mux-Smm/artist/controllers"
	"Go-Mux-Smm/artist/repositories"
	"Go-Mux-Smm/artist/services"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main () {
	port := "9092"

	connect := "root:@tcp(127.0.0.1:3306)/kotlify"

	db, err := sql.Open("mysql", connect)
	if err != nil {
		log.Fatal("Error When Connect to DB " + connect + " : " + err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	var artistRepo artist.ArtistRepo  = repositories.CreateArtistRepoMysqlImpl(db)
	var artistService = services.CreateArtistService(artistRepo)
	controllers.CreateArtistRoutes(router, artistService)


	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
