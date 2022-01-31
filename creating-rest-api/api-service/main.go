package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.GET("/shows", getShows)
	router.POST("/shows", createShows)

	router.Run("localhost:8081")
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