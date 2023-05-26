package utils

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetJSON(url string) (gjson.Result, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return gjson.Result{}, err
	}

	SetReqHeaders(req)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return gjson.Result{}, err
	}
	rateLimitRemaining := res.Header.Get("X-RateLimit-UserRemaining")
	if rateLimitRemaining != "" {
		ratelimit, _ := strconv.Atoi(rateLimitRemaining)
		if ratelimit <= 0 {
			return gjson.Result{}, fmt.Errorf("ratelimited by imgur")
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return gjson.Result{}, err
	}

	switch (res.StatusCode) {
	case 200:
		return gjson.Parse(string(body)), nil
	case 429:
		return gjson.Result{}, fmt.Errorf("ratelimited by imgur")
	default:
		return gjson.Result{}, fmt.Errorf("received status %s, expected 200 OK.\n%s", res.Status, string(body))
	}
}