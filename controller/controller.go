package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"songs/apperrors"
	"songs/config"
	"songs/song"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// @Summary		Create song
// @Description	Find info by song name and group then record it
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param dto body CreateSongDTO true "Data for create song"
// @Success		201	{object}	song.Song
// @Failure     400 "bad request data"
// @Router			/songs [post]
func CreateSong(c *gin.Context) {
	var dto CreateSongDTO
	err := validateDTO(c, &dto)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	client := resty.New()
	resp, _ := client.R().
		SetQueryParam("song", dto.Song).
		SetQueryParam("group", dto.Group).
		Get(config.InfoApiUrl)

	var respBody ApiResponse
	err = json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(apperrors.ErrBadApiResponse))
		return
	}

	newSong, err := song.Create(
		dto.Song,
		dto.Group,
		respBody.ReleaseDate,
		respBody.Text,
		respBody.Link,
	)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	c.IndentedJSON(http.StatusCreated, newSong)
}

// @Summary		Find song
// @Description	Find by info and paginate
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param dto body FindSongDTO true "Data for find song"
// @Success		200	{object}	[]song.Song
// @Failure     400 "bad request data"
// @Router			/songs/find [post]
func FindSongs(c *gin.Context) {
	var dto FindSongDTO
	err := validateDTO(c, &dto)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	var findSongErrors []error

	offset := transformToInt(dto.Offset, "offset", &findSongErrors)
	limit := transformToInt(dto.Limit, "limit", &findSongErrors)

	dateBegin := transformToInt64(
		dto.ReleaseDateBegin,
		"releaseDateBegin",
		&findSongErrors,
	)

	dateEnd := transformToInt64(
		dto.ReleaseDateEnd,
		"releaseDateEnd",
		&findSongErrors,
	)

	if len(findSongErrors) != 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			fmt.Sprintf("Validation error: %s", errors.Join(findSongErrors...)),
		)
		return
	}

	songs, err := song.Find(
		dto.Song, dto.Group, dto.Text, dto.Link, dateBegin, dateEnd,
		offset, limit,
	)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	c.IndentedJSON(http.StatusOK, songs)
}

// @Summary		Get text
// @Description	Get song text and paginate by verses
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param dto body GetTextDTO true "Data for get song text"
// @Success		200	{object}	[]string
// @Failure     400 "bad request data"
// @Router			/songs/text [post]
func GetText(c *gin.Context) {
	var dto GetTextDTO
	err := validateDTO(c, &dto)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	var findSongErrors []error

	offset := transformToInt(dto.Offset, "offset", &findSongErrors)
	limit := transformToInt(dto.Limit, "limit", &findSongErrors)

	verses, err := song.GetText(dto.ID, int(offset), int(limit))
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	c.IndentedJSON(http.StatusOK, verses)
}

// @Summary		Delete song
// @Description	Delete song by id
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param dto body IdDTO true "Data with song id"
// @Success		200	{object}	song.Song
// @Failure     400 "bad request data"
// @Router			/songs [delete]
func DeleteSong(c *gin.Context) {
	var dto IdDTO
	err := validateDTO(c, &dto)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}
	deletedSong, err := song.Delete(dto.ID)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	c.IndentedJSON(http.StatusOK, deletedSong)
}

// @Summary		Update song
// @Description	Update song info
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param dto body UpdateSongDTO true "Data for update song info"
// @Success		200	{object}	song.Song
// @Failure     400 "bad request data"
// @Router			/songs [put]
func UpdateSong(c *gin.Context) {
	var dto UpdateSongDTO
	err := validateDTO(c, &dto)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	newSong, err := song.Update(
		dto.ID,
		dto.Song,
		dto.Group,
		dto.ReleaseDate,
		dto.Text,
		dto.Link,
	)
	if err != nil {
		c.IndentedJSON(apperrors.GetErrorAndStatus(err))
		return
	}

	c.IndentedJSON(http.StatusOK, newSong)
}
