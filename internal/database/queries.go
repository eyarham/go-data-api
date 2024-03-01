package queries

import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var table = "album"

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumByID queries for the album with the specified ID.
func AlbumByID(id int64) (Album, error) {
	openConnection()

	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM "+table+" WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func AllAlbums() ([]Album, error) {
	openConnection()
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q", err)
	}
	return albums, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func Add(alb Album) (int64, error) {
	openConnection()
	result, err := db.Exec("INSERT INTO "+table+" (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func openConnection() {
	cfg := mysql.Config{
		User:   "apiservice",  // os.Getenv("DBUSER"),
		Passwd: "badpassword", //os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}