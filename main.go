package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/artist/:name", getAlbumByArtist)
	router.POST("/albums", postAlbums)
	router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	// Can be substituted with c.JSON() to pass compact JSON
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusCreated, gin.H{"Message": "ID already taken, added to last one"})
			return
		}
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found"})
}

func getAlbumByArtist(c *gin.Context) {
	var albumsArtist []album
	artist := c.Param("name")
	artist = strings.Replace(artist, "+", " ", -1)

	for _, a := range albums {
		if a.Artist == artist {
			albumsArtist = append(albumsArtist, a)
		}
	}

	if albumsArtist != nil {
		c.IndentedJSON(http.StatusOK, albumsArtist)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Artist not found"})
}
