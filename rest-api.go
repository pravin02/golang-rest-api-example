package main

import "net/http"
import "github.com/gin-gonic/gin"

type album struct {
	
	ID string `json:"albumId"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
	
}

var albums = []album {}

func getAllAlbums(c *gin.Context) {
	if len(albums) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "No albums found"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)	
}

func addAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("albums/:id", getAlbumByID)
	router.POST("/albums", addAlbum)
	
	router.Run("localhost:8080")

}