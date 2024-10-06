package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func GetInfo(c *gin.Context) {
	song := c.Query("song")
	if song == "" {
		c.IndentedJSON(http.StatusBadRequest, "param name not found")
		return
	}

	group := c.Query("group")
	if group == "" {
		c.IndentedJSON(http.StatusBadRequest, "param group not found")
		return
	}

	beginDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	endDate := time.Now().Unix()

	randomDate := rand.Int63n(endDate-beginDate) + beginDate

	var sb strings.Builder

	for k := 0; k < 4+rand.Intn(4); k++ {
		for j := 0; j < 2+rand.Intn(6); j++ {
			for i := 0; i < 30+rand.Intn(10); i++ {
				sb.WriteRune(rune('a' + rand.Intn(26)))
			}
			sb.WriteRune('\n')
		}
		sb.WriteRune('\n')
	}

	c.IndentedJSON(http.StatusOK, ApiResponse{
		ReleaseDate: time.Unix(randomDate, 0).Format("02.01.2006"),
		Text:        sb.String(),
		Link:        fmt.Sprintf("https://www.somesite.com/%v/%v", group, song),
	})
}
