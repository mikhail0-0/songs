package song

import (
	"songs/apperrors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Song struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id" example:"147367f5-93ef-432d-8a97-b06f716f9fad"`

	Name        string    `gorm:"type:varchar" json:"name" example:"Supermassive Black Hole"`
	Group       string    `gorm:"type:varchar" json:"group" example:"Muse"`
	ReleaseDate time.Time `gorm:"type:timestamptz" json:"releaseDate" example:"2006-07-16T00:00:00+03:00"`
	Text        string    `gorm:"type:text" json:"text" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link        string    `gorm:"type:varchar" json:"link" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

var db *gorm.DB

func Init(refDB *gorm.DB) {
	db = refDB
	if !db.Migrator().HasTable(&Song{}) {
		db.Migrator().CreateTable(&Song{})
	}
}

func Create(songName, groupName, releaseDateStr, text, link string) (*Song, error) {
	song := Song{
		Name:  songName,
		Group: groupName,
		Text:  text,
		Link:  link,
	}

	releaseDate, err := timeFromString(releaseDateStr)
	if err != nil {
		return nil, err
	}
	song.ReleaseDate = *releaseDate

	result := db.Create(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	return &song, nil
}

func Update(strId, songName, groupName, releaseDateStr, text, link string) (*Song, error) {
	id, err := uuid.Parse(strId)
	if err != nil {
		return nil, err
	}

	song := Song{
		ID:    id,
		Name:  songName,
		Group: groupName,
		Text:  text,
		Link:  link,
	}

	if releaseDateStr != "" {
		releaseDate, err := timeFromString(releaseDateStr)
		if err != nil {
			return nil, err
		}
		song.ReleaseDate = *releaseDate
	}

	result := db.Model(&song).Clauses(clause.Returning{}).Updates(&song)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, apperrors.ErrNotFound
	}

	return &song, nil
}

func Find(
	songName, groupName, text, link string,
	releaseDateBegin, releaseDateEnd int64,
	offset, limit int,
) ([]Song, error) {
	expressions := make([]clause.Expression, 0)
	if songName != "" {
		expressions = append(expressions, clause.Like{Column: "name", Value: likeStr(songName)})
	}
	if groupName != "" {
		expressions = append(expressions, clause.Like{Column: "group", Value: likeStr(groupName)})
	}
	if text != "" {
		expressions = append(expressions, clause.Like{Column: "text", Value: likeStr(text)})
	}
	if link != "" {
		expressions = append(expressions, clause.Like{Column: "link", Value: likeStr(link)})
	}

	if releaseDateBegin != 0 && releaseDateEnd != 0 {
		expressions = append(expressions, clause.NamedExpr{
			SQL: "release_date BETWEEN ? AND ?",
			Vars: []any{
				time.Unix(releaseDateBegin, 0),
				time.Unix(releaseDateEnd, 0),
			},
		})
	}

	var refLim *int
	if limit == 0 {
		refLim = nil
	} else {
		refLim = &limit
	}

	expressions = append(expressions, clause.Limit{Offset: offset, Limit: refLim})

	var songs []Song
	result := db.Clauses(expressions...).Find(&songs)

	if result.Error != nil {
		return nil, result.Error
	}

	return songs, nil
}

func GetText(songId string, offset, limit int) ([]string, error) {
	var song Song
	result := db.First(&song, "id = ?", songId)
	if result.Error != nil {
		return nil, apperrors.ErrNotFound
	}

	verses := strings.Split(song.Text, "\n\n")

	var end int
	if limit == 0 {
		end = len(verses)
	} else {
		end = offset + limit
	}

	return verses[offset:end], nil
}

func Delete(songId string) (*Song, error) {
	var song Song

	result := db.Model(&Song{}).Clauses(clause.Returning{}).Delete(&song, "id = ?", songId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, apperrors.ErrNotFound
	}

	return &song, nil
}
