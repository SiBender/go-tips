package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	//connection
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "qwe123"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	//get database handle
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

	res, err := albumsByArtists("John Coltrane\" or 1 = 1 -- ")
	fmt.Println(res)

	alb, err := albumByID(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(alb)
	/*
		albID, err := addAlbum(Album{
			Title:  "Brothers In Arms",
			Artist: "Dire Straits",
			Price:  49.99,
		})

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ID of added album: %v\n", albID)*/
}

func albumsByArtists(name string) ([]Album, error) {
	//Album slice
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtists: %q, %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtists: %q, %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtists: %q, %v", name, err)
	}

	return albums, nil
}

func albumByID(id int) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID: %d: %v", id, err)
	}

	return alb, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}
