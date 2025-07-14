package main

// import "net/http"
import "github.com/gin-gonic/gin"

func getAllAlbums(c *gin.Context) {
}

func getAlbumByID(c *gin.Context) {
}

func addAlbum(c *gin.Context) {
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("albums/:id", getAlbumByID)
	router.POST("/albums", addAlbum)
	
	router.Run("localhost:8080")

}