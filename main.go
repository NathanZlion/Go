package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID	 string `json:"id"`
	Title  string  `json:"title"`
    Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type albumUpdate struct {
	Title  string  `json:"title"`
    Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums []album = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums);
}

func postAlbum(context *gin.Context) {
	var newAlbum album

	err := context.BindJSON(&newAlbum);
	if err != nil {
		return
	}

	albums = append(albums, newAlbum);

	context.IndentedJSON(http.StatusOK, newAlbum);
}

func getAlbumById(context *gin.Context) {
	id := context.Param("id");

	for _, album := range(albums) {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album);
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, map[string]any {"message": "album with that id doesn't exist."});
}

func updateAlbum(context *gin.Context) {
	id := context.Param("id");
	var update albumUpdate

	if err := context.BindJSON(&update); err != nil {
		return 
	}

	for index, album := range albums {
		if album.ID == id {
			albums[index].Title = update.Title;
			albums[index].Price = update.Price;
			albums[index].Artist = update.Artist;
			context.IndentedJSON(http.StatusOK, gin.H{"message": "update successfull"})
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cannot find an album with that id"})
}

func deleteAtIndex(s []album, i int) []album {
    return append(s[:i], s[i+1:]...)
}

func deleteAlbum(context *gin.Context) {
	id := context.Param("id");

	for index, album := range albums {
		if album.ID == id {
			albums = deleteAtIndex(albums, index);
			context.IndentedJSON(http.StatusOK, gin.H{"message": "Delete Successfull"});
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cannot find an album with that id"})
}

func main() {
	router := gin.Default();
	router.GET("/albums", getAlbums);
	router.POST("/albums", postAlbum);
	router.GET("/albums/:id", getAlbumById);
	router.PATCH("/albums/:id", updateAlbum);
	router.DELETE("/albums/:id", deleteAlbum);
	router.Run("localhost:8080");
}