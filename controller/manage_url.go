package controller

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/EmeraldLS/url-shortener/model"
	"github.com/EmeraldLS/url-shortener/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-module/carbon"
)

// It shortens the url with a fixed expiration time of 5mins
func ShortenURL(c *gin.Context) {
	myURL := c.Request.URL

	var req model.URL_SHORTENER_REQUEST
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": fmt.Sprintf("binding error -> %v", err),
		})
		c.Abort()
		return
	}

	if err := validator.New().Struct(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": fmt.Sprintf("validation error -> %v", err),
		})
		c.Abort()
		return
	}

	var url = new(model.URL)

	url.OriginalUrl = req.OriginalUrl
	url.Id = Random_id()
	url.ShortUrl = myURL.Hostname() + "/" + url.Id
	url.CreatedAt = carbon.Now().ToDateTimeString()
	url.UpdatedAt = carbon.Now().ToDateTimeString()
	url.ExpirationTime = time.Now().Add(time.Minute * 5)
	if err := service.SaveUrl(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": fmt.Sprintf("saving error -> %v", err),
		})
	}

	c.JSON(http.StatusCreated, url)
}

func generateRandomNumber(length int) int {
	rand.NewSource(time.Now().UnixNano())

	return rand.Intn(length)
}
func Random_id() string {
	var needed []string
	var letters = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for {
		letter := letters[generateRandomNumber(len(letters))]
		needed = append(needed, string(letter))
		if len(needed) == 6 {
			break
		}
	} //time complexity 0(n)

	id := strings.Join(needed, "")
	return id
}

func RouteToURL(c *gin.Context) {
	id := c.Param("id")
	url_content, err := service.GetURlId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	go func(url_id string) {
		time.Sleep(time.Second * 5)
		err = service.UpdateURLClicks(url_id)
		if err != nil {
			log.Print(err)
		}
	}(id)

	if url_content.OriginalUrl[0:8] == "https://" || url_content.OriginalUrl[0:7] == "http://" {
		log.Println("Idc")
	} else {
		log.Println("This was evaluated")
		url_content.OriginalUrl = "https://" + url_content.OriginalUrl
	}

	c.Redirect(http.StatusFound, url_content.OriginalUrl)
}

func DeleteExpiredURLS(c *gin.Context) {
	if err := service.DeleteExpiredURLS(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "expired urls deleted",
	})
}
