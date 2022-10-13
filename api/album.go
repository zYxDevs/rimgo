package api

import (
	"strings"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/microcosm-cc/bluemonday"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
)

type Album struct {
	Id                  string
	Title               string
	Views               int64
	Upvotes             int64
	Downvotes           int64
	SharedWithCommunity bool
	CreatedAt           string
	UpdatedAt           string
	Comments            int64
	User                User
	Media               []Media
	Tags                []Tag
}

type Media struct {
	Id          string
	Name        string
	Title       string
	Description string
	Url         string
	Type        string
	MimeType    string
}

var albumCache = cache.New(1*time.Hour, 15*time.Minute)

func FetchAlbum(albumID string) (Album, error) {
	cacheData, found := albumCache.Get(albumID + "-album")
	if found {
		return cacheData.(Album), nil
	}

	data, err := utils.GetJSON("https://api.imgur.com/post/v1/albums/" + albumID + "?client_id=" + utils.Config.ImgurId + "&include=media%2Caccount")
	if err != nil {
		return Album{}, err
	}

	album, err := ParseAlbum(data)
	if err != nil {
		return Album{}, err
	}

	albumCache.Set(albumID+"-album", album, cache.DefaultExpiration)
	return album, err
}

func FetchPosts(albumID string) (Album, error) {
	cacheData, found := albumCache.Get(albumID + "-posts")
	if found {
		return cacheData.(Album), nil
	}

	data, err := utils.GetJSON("https://api.imgur.com/post/v1/posts/" + albumID + "?client_id=" + utils.Config.ImgurId + "&include=media%2Caccount%2Ctags")
	if err != nil {
		return Album{}, err
	}

	album, err := ParseAlbum(data)
	if err != nil {
		return Album{}, err
	}

	albumCache.Set(albumID+"-posts", album, cache.DefaultExpiration)
	return album, nil
}

func FetchMedia(mediaID string) (Album, error) {
	cacheData, found := albumCache.Get(mediaID + "-media")
	if found {
		return cacheData.(Album), nil
	}

	data, err := utils.GetJSON("https://api.imgur.com/post/v1/media/" + mediaID + "?client_id=" + utils.Config.ImgurId + "&include=media%2Caccount")
	if err != nil {
		return Album{}, err
	}

	album, err := ParseAlbum(data)
	if err != nil {
		return Album{}, err
	}

	albumCache.Set(mediaID+"-media", album, cache.DefaultExpiration)
	return album, nil
}

func ParseAlbum(data gjson.Result) (Album, error) {
	media := make([]Media, 0)
	data.Get("media").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			url := value.Get("url").String()
			url = strings.ReplaceAll(url, "https://i.imgur.com", "")

			description := value.Get("metadata.description").String()
			description = strings.ReplaceAll(description, "\n", "<br>")
			description = bluemonday.UGCPolicy().Sanitize(description)

			media = append(media, Media{
				Id:          value.Get("id").String(),
				Name:        value.Get("name").String(),
				MimeType:    value.Get("mime_type").String(),
				Type:        value.Get("type").String(),
				Title:       value.Get("metadata.title").String(),
				Description: description,
				Url:         url,
			})

			return true
		},
	)

	tags := make([]Tag, 0)
	data.Get("tags").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			tags = append(tags, Tag{
				Tag:          value.Get("tag").String(),
				Display:      value.Get("display").String(),
				Background:   "/" + value.Get("background_id").String() + ".webp",
				BackgroundId: value.Get("background_id").String(),
			})
			return true
		},
	)

	createdAt, err := utils.FormatDate(data.Get("created_at").String())
	if err != nil {
		return Album{}, err
	}

	album := Album{
		Id:                  data.Get("id").String(),
		Title:               data.Get("title").String(),
		SharedWithCommunity: data.Get("shared_with_community").Bool(),
		Views:               data.Get("view_count").Int(),
		Upvotes:             data.Get("upvote_count").Int(),
		Downvotes:           data.Get("downvote_count").Int(),
		Comments:            data.Get("comment_count").Int(),
		CreatedAt:           createdAt,
		Media:               media,
		Tags:                tags,
	}

	account := data.Get("account")
	if account.Raw != "" {
		album.User = User{
			Id:       account.Get("id").Int(),
			Username: account.Get("username").String(),
			Avatar:   strings.ReplaceAll(account.Get("avatar_url").String(), "https://i.imgur.com", ""),
		}
	}

	return album, nil
}
