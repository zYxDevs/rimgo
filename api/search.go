package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"codeberg.org/rimgo/rimgo/utils"
	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Id        string
	Url				string
	ImageUrl	string
	Title			string
	User			string
	Points	  string
	Views			string
	RelTime		string
}

func (client *Client) Search(query string) ([]SearchResult, error) {
	query = url.QueryEscape(query)
	req, err := http.NewRequest("GET", "https://imgur.com/search?qs=list&q=" + query, nil)
	if err != nil {
		return []SearchResult{}, err
	}
	utils.SetReqHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []SearchResult{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
    return []SearchResult{}, fmt.Errorf("invalid status code, got %d", res.StatusCode)
  }

	doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    return []SearchResult{}, err
  }

	results := []SearchResult{}
  doc.Find(".post-list").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		imageUrl, _ := s.Find("img").Attr("src")

		postInfo := strings.Split(s.Find(".post-info").Text(), "·")
		points := strings.TrimSpace(postInfo[0])
		points = strings.TrimSuffix(points, " points")
		views := strings.TrimSpace(postInfo[1])
		views = strings.TrimSuffix(views, " views")

		result := SearchResult{
			Id: strings.Split(url, "/")[2],
			Url: url,
			ImageUrl: strings.ReplaceAll(imageUrl, "//i.imgur.com", ""),
			Title: s.Find(".search-item-title a").Text(),
			User: s.Find(".account").Text(),
			Views: views,
			Points: points,
			RelTime: strings.TrimSpace(postInfo[2]),
		}

		results = append(results, result)
	})

	return results, nil
}