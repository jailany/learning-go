package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type show struct {
	ID	  int 	 `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Likes int    `json:"likes"`
}

var shows = []show{
	{
		ID: 1,
		Title: "dark",
		Genre: "thriller",
		Likes: 200,
	},
	{
		ID: 2,
		Title: "the office",
		Genre: "comedy",
		Likes: 400,
	},
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/shows", getShows)
	router.GET("/shows/:id", getShowByID)
	router.POST("/shows", createShows)

	router.Run("localhost:8081")
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func getShows(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, shows)
}

func createShows(c *gin.Context) {
	var newShow show

	if err := c.BindJSON(&newShow); err != nil {
		return
	}

	shows = append(shows, newShow)
	c.IndentedJSON(http.StatusCreated, shows);
}

func getShowByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range shows {
		convertedId, err := strconv.Atoi(id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		}
		if a.ID == convertedId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "show not found"})
}