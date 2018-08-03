package server

import (
	"encoding/json"
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
	TOTAL_POEMS   = 2300
	TOTAL_ROBAEIS = 1992
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
	router.POST("/", indexHandler)
	router.POST("/random/:type/", randomHandler)
	router.POST("/ghazal/:number/", showHandler)
	router.POST("/search/", searchHandler)

	router.POST("/favorites", underConstructionHandler)
	router.POST("/mostviewed", underConstructionHandler)

	return router.Run("127.0.0.1:8081")
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", nil)
}

func searchHandler(ctx *gin.Context) {
	payload := ctx.PostForm("payload")
	var Payload struct {
		Word string `json:"word"`
	}
	json.Unmarshal([]byte(payload), &Payload)

	poems, err := database.FindPoem(Payload.Word)
	if err != nil {
		log.Println("Error on searching", Payload.Word, err)
		ctx.HTML(http.StatusOK, "under-construction", gin.H{"message": "خطایی رخ داده است"})
		return
	}
	ctx.HTML(http.StatusOK, "search-result", gin.H{"poems": poems})
}

func randomHandler(ctx *gin.Context) {
	mode := ctx.Param("type")
	switch mode {
	case "ghazal":
		showPoem(ctx, 1+rand.Intn(TOTAL_POEMS))
		break
	case "robaei":
		showRobaei(ctx, 1+rand.Intn(TOTAL_ROBAEIS))
		break
	}
}

func showHandler(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.HTML(http.StatusOK, "under-construction", gin.H{"message": "خطایی رخ داده است"})
		return
	}
	showPoem(ctx, number)
}

func showPoem(ctx *gin.Context, number int) {
	poem, err := database.GetPoem("ID", number)
	if err != nil {
		log.Println("Error on fetching [ghazal]", number, err)
		ctx.HTML(http.StatusOK, "under-construction", gin.H{"message": "خطایی رخ داده است"})
		return
	}
	ctx.HTML(http.StatusOK, "show", gin.H{
		"number":  number,
		"content": poem.Content,
		"type":    1,
	})
}

func showRobaei(ctx *gin.Context, number int) {
	robaei, err := database.GetRobaei("ID", number)
	if err != nil {
		log.Println("Error on fetching [robaei]", number, err)
		ctx.HTML(http.StatusOK, "under-construction", gin.H{"message": "خطایی رخ داده است"})
		return
	}
	ctx.HTML(http.StatusOK, "show", gin.H{
		"number":  number,
		"content": robaei.Content,
		"type":    2,
	})
}

func underConstructionHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "under-construction", nil)
}
