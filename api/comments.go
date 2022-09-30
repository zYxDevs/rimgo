package api

import (
	"strings"
	"sync"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/dustin/go-humanize"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
)

type Comment struct {
	Comments  []Comment
	User			User
	Id        string
	Comment   string
	Upvotes   int64
	Downvotes int64
	Platform  string
	CreatedAt string
	RelTime		string
	UpdatedAt string
	DeletedAt string
}

var commentCache = cache.New(15*time.Minute, 15*time.Minute)

func FetchComments(galleryID string) ([]Comment, error) {
	cacheData, found := commentCache.Get(galleryID)
	if found {
		return cacheData.([]Comment), nil
	}

	data, err := utils.GetJSON("https://api.imgur.com/comment/v1/comments?client_id=" + utils.Config.ImgurId + "&filter[post]=eq:" + galleryID + "&include=account,adconfig&per_page=30&sort=best")
	if err != nil {
		return []Comment{}, nil
	}

	wg := sync.WaitGroup{}
	comments := make([]Comment, 0)
	data.Get("data").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				comments = append(comments, ParseComment(value))
			}()

			return true
		},
	)
	wg.Wait()

	commentCache.Set(galleryID, comments, cache.DefaultExpiration)
	return comments, nil
}

func ParseComment(data gjson.Result) Comment {
	createdTime, _ := time.Parse("2006-01-02T15:04:05Z", data.Get("created_at").String())
	createdAt := createdTime.Format("January 2, 2006 3:04 PM")
	updatedAt, _ := utils.FormatDate(data.Get("updated_at").String())
	deletedAt, _ := utils.FormatDate(data.Get("deleted_at").String())

	userAvatar := strings.ReplaceAll(data.Get("account.avatar").String(), "https://i.imgur.com", "")

	wg := sync.WaitGroup{}
	comments := make([]Comment, 0)
	data.Get("comments").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				comments = append(comments, ParseComment(value))
			}()

			return true
		},
	)
	wg.Wait()

	return Comment{
		Comments: comments,
		User: User{
			Id:       data.Get("account.id").Int(),
			Username: data.Get("account.username").String(),
			Avatar:   userAvatar,
		},
		Id:        data.Get("id").String(),
		Comment:   data.Get("comment").String(),
		Upvotes:   data.Get("upvote_count").Int(),
		Downvotes: data.Get("downvote_count").Int(),
		Platform:  data.Get("platform").String(),
		CreatedAt: createdAt,
		RelTime: 	 humanize.Time(createdTime),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
