package main

import (
	queries "github.com/eyarham/go-data-api/internal/database"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albumsFromDb, err := queries.AllAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, albumsFromDb)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	var dbAlbum queries.Album
	i, err := queries.Add(dbAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	newAlbum.ID = strconv.FormatInt(i, 10)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// ... handle error
		panic(err)
	}
	album, err := queries.AlbumByID(i)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found" + err.Error()})
	}
	c.IndentedJSON(http.StatusOK, album)
}
