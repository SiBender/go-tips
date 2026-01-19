package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbum)

	router.Run("localhost:8080")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = map[int]album{
	1: {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	2: {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	3: {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(context *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := context.BindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Add new
	id, _ := strconv.Atoi(newAlbum.ID)
	albums[id] = newAlbum
	fmt.Println(len(albums))
	fmt.Println(albums)

	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	found, exists := albums[id]
	if exists {
		context.IndentedJSON(http.StatusOK, found)
	} else {
		context.IndentedJSON(http.StatusOK, gin.H{"message": "album not found"})
	}
}
