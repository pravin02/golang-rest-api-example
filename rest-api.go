package main

import "net/http"
import "github.com/gin-gonic/gin"


// defining album dto for rest api
type album struct {
	ID string `json:"albumId"` // ID mapped with albumId json key
	Title string `json:"title"`  // Title mapped with title json key
	Artist string `json:"artist"` // Artist mapped with artist json key
	Price float64 `json:"price"` // Price mapped with price json key
}

// defined albums empty array
var albums = []album {}

// defined function to return all the albums
func getAllAlbums(c *gin.Context) {
	if len(albums) == 0 { // if albums array empty then returning 404 Not Found
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No albums found"})
		return
	}
	// returning albums
	c.IndentedJSON(http.StatusOK, albums)
}

// get album by album id
func getAlbumByID(c *gin.Context) {
	albumId := c.Param("id") // extract the id from route
	
	// loop the existing albums array and find the albumId matching with arrays ID field
	for _, a := range albums {
		if a.ID == albumId {
			c.IndentedJSON(http.StatusOK, a) // return the album object if matching id found
			return // breking the loop and returning from the function
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No album found by Id "+albumId})	
}

// below method take gin context as an input to bind the request payload json into newAlbum method
func addAlbum(c *gin.Context) {
	var newAlbum album
	
	if err := c.BindJSON(&newAlbum); err != nil {
		return // if error occurred while binding json into newAlbum field as per album data type
	}	
	albums = append(albums, newAlbum) // newAlbum object into albums array and update the albums variable with new array
	c.IndentedJSON(http.StatusCreated, newAlbum) // return newAlbum object with 201 status code
}

func main() {
	router := gin.Default() // Creating the router
	router.GET("/albums", getAllAlbums) // get route to return the albums list
	router.GET("albums/:id", getAlbumByID) // get route to return the album if found in albums if not then with error message with 404
	router.POST("/albums", addAlbum) // add new album in albums array
	
	router.Run("localhost:8080") // run the server on 8080 port with routes configurations
}


/*

run below commands on git bash
curl http://localhost:8080/albums \
	--header "Content-Type: application/json" \
	--request "GET"

curl http://localhost:8080/albums/1 \
	--header "Content-Type: application/json" \
	--request "GET"

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

curl http://localhost:8080/albums/4 \
	--header "Content-Type: application/json" \
	--request "GET"

*/