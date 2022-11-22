package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/rs/xid"
	"net/http"
	"io/ioutil"
	"encoding/json"
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

func DeleteRecipeHandler(context *gin.Context) {
	id:=context.Param("id")
	index:=-1

	for i:=0; i<len(recipes); i++ {
		if recipes[i].ID == id {
			index=i
		}
	}

	if index == -1{
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...)

	context.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted"})
}

func UpdateRecipeHandler(context *gin.Context) {
	id := context.Param("id")
	var recipe Recipe

	if err := context.ShouldBindJSON(&recipe); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index=i
		}
	}

	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}

	recipes[index]=recipe
	context.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(context *gin.Context) {
	context.JSON(http.StatusOK, recipes)
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

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.Run()
}