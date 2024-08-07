package v1

import (
	"context"
	"net/http"

	"github.com/Vghxv/GinHub/database"
	"github.com/Vghxv/GinHub/models"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, title, artist, price FROM albums")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer rows.Close()

	albums := []models.Album{}
	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		albums = append(albums, a)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var a models.Album
	err := database.DB.QueryRow(context.Background(), "SELECT id, title, artist, price FROM albums WHERE id=$1", id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, a)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	err := database.DB.QueryRow(context.Background(), "INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id", newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&newAlbum.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}