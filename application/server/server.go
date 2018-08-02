package server

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ahmdrz/divan-e-shams/application/server/template"
	"github.com/ahmdrz/divan-e-shams/database"
)

const (
	TOTAL_POEMS = 2300
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run() error {
	err := database.Open("database.boltdb")
	if err != nil {
		return err
	}

	router := gin.Default()
	router.SetHTMLTemplate(template.New("./templates"))

	router.Static("/resources", "./resources")
	router.GET("/", indexHandler)
	router.GET("/random/", randomHandler)
	router.GET("/ghazal/:number/", showHandler)

	return router.Run("127.0.0.1:8081")
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", nil)
}

func randomHandler(ctx *gin.Context) {
	randomNumber := 1 + rand.Intn(TOTAL_POEMS)
	showPoem(ctx, randomNumber)
}

func showHandler(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		log.Println("Input error", err)
		return
	}
	showPoem(ctx, number)
}

func showPoem(ctx *gin.Context, number int) {
	poem, err := database.GetPoem("ID", number)
	if err != nil {
		log.Println("Database error", err)
		return
	}
	ctx.HTML(http.StatusOK, "show", gin.H{
		"number":  number,
		"content": poem.Content,
	})
}
