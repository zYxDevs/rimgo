package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"codeberg.org/rimgo/rimgo/utils"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
)

func (client *Client) FetchTrending(section, sort, page string) ([]Submission, error) {
	cacheData, found := client.Cache.Get(fmt.Sprintf("trending-%s-%s-%s", section, sort, page))
	if found {
		return cacheData.([]Submission), nil
	}

	req, err := http.NewRequest("GET", "https://api.imgur.com/post/v1/posts", nil)
	if err != nil {
		return []Submission{}, err
	}
	utils.SetReqHeaders(req)

	q := req.URL.Query()
	q.Add("client_id", client.ClientID)
	q.Add("include", "cover")
	q.Add("page", page)

	switch sort {
	case "newest":
		q.Add("filter[window]", "week")
		q.Add("sort", "-time")
	case "best":
		q.Add("filter[window]", "all")
		q.Add("sort", "-top")
	case "popular":
		fallthrough
	default:
		q.Add("filter[window]", "week")
		q.Add("sort", "-viral")
		sort = "popular"
	}
	switch section {
	case "hot":
		q.Add("filter[section]", "eq:hot")
	case "new":
		q.Add("filter[section]", "eq:new")
	case "top":
		q.Add("filter[section]", "eq:top")
		q.Add("filter[window]", "day")
	default:
		q.Add("filter[section]", "eq:hot")
		section = "hot"
	}

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Submission{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Submission{}, err
	}

	data := gjson.Parse(string(body))

	wg := sync.WaitGroup{}
	posts := make([]Submission, 0)
	data.ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				posts = append(posts, Submission{
					Id:    value.Get("id").String(),
					Title: value.Get("title").String(),
					Link:  strings.ReplaceAll(value.Get("url").String(), "https://imgur.com", ""),
					Cover: Media{
						Id:   value.Get("cover_id").String(),
						Type: value.Get("cover.type").String(),
						Url:  strings.ReplaceAll(value.Get("cover.url").String(), "https://i.imgur.com", ""),
					},
					Points:    value.Get("point_count").Int(),
					Upvotes:   value.Get("upvote_count").Int(),
					Downvotes: value.Get("downvote_count").Int(),
					Comments:  value.Get("comment_count").Int(),
					Views:     value.Get("view_count").Int(),
					IsAlbum:   value.Get("is_album").Bool(),
				})
			}()

			return true
		},
	)

	wg.Wait()

	client.Cache.Set(fmt.Sprintf("trending-%s-%s-%s", section, sort, page), posts, cache.DefaultExpiration)
	return posts, nil
}
