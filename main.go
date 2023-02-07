package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"mvc/app/controllers"
	"mvc/app/models"
	"mvc/app/views"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   `root`,
		Passwd: `password`,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my-app",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	_, err = db.Exec("CREATE table if not exists artist (id int AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), painting VARCHAR(255), invited BOOLEAN)")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS exhibition (id INT AUTO_INCREMENT PRIMARY KEY, artist_id INT, FOREIGN KEY (artist_id) REFERENCES artist(id))")
	if err != nil {
		panic(err.Error())
	}

	var eid int
	ac := controllers.ArtistController{}
	ec := controllers.ExhibitionController{}

Text:
	fmt.Printf("CHOOSE WHAT iaTO DO: " +
		"\nInvite artist: ia" +
		"\nOrganize exhibition: oe" +
		"\nAdd artist's paintings to the gallery: ap" +
		"\nDelete artist's paintings from the gallery: dp" +
		"\nExit: e\n")
	var command string
	_, _ = fmt.Scan(&command)

	switch command {
	case `ia`:
		a := &models.Artist{`Roman`, `black box`, true, 23, "neolideralizm"}
		ac.GetALLArtists()
		fmt.Println(" ----- Artist invited")
		if err := views.PrintArtist(a); err != nil {
			fmt.Println(err)
		}
		goto Text
	case `oe`:
		eid++
		e := models.NewExhibition(eid)
		ec.CreateExhibition(e)
		fmt.Println(" ----- Exhibition opened")
		if err := views.PrintExhibition(e); err != nil {
			fmt.Println(err)
		}
		goto Text
	case `ap`:
		fmt.Printf("<artist's painting>?")
		var artistPainting string
		_, _ = fmt.Scan(&artistPainting)
		artist2add := ac.DeleteArtist(artistPainting)
		if artist2add == nil {
			fmt.Printf("Painting of %s not found", artistPainting)
			goto Text
		}
		fmt.Printf("<exhibition id>?")
		var exhibitionId int
		_, _ = fmt.Scan(&exhibitionId)
		exhibition := ec.AddArtist(exhibitionId, artist2add)
		fmt.Println(" ----- Artist added")
		if err := views.PrintExhibition(exhibition); err != nil {
			fmt.Println(err)
		}
		goto Text
	case `dp`:
		fmt.Printf("<Artist's painting>?")
		var artistPainting string
		_, _ = fmt.Scan(&artistPainting)
		artist2delete := ac.DeleteArtist(artistPainting)
		if artist2delete == nil {
			fmt.Printf("User with %s not found", artistPainting)
			goto Text
		}
		fmt.Printf("<course id>?")
		var exhibitionId int
		_, _ = fmt.Scan(&exhibitionId)
		ec.DeleteArtist(exhibitionId, artist2delete)
		fmt.Println(" ----- Painting taken off the exhibition")
		goto Text
	case `e`:
		break
	}

	fmt.Println("Closed by hoster")
}
