package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/rs/xid"
	"net/http"
)

var recipes []Recipe

type Recipe struct {
	ID         string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func NewRecipeHandler(context *gin.Context) {
	var recipe Recipe

	if err := context.ShouldBindJSON(&recipe); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	context.JSON(http.StatusOK, recipe)
}

func init(){
	recipes = make([]Recipe, 0)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run()
}