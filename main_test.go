package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	endpoint, err := url.Parse("https://hackersandslackers.app/members/api/member/")
	cookie := "__stripe_mid=939f07db-3218-4e38-b192-4616eef64ec0; em_cdn_uid=t%3D1587926226095%26u%3Dde7c8991dea14a95b1a1c0036b9c6372; em_p_uid=l:1592880910636|t:1587926226510|u:b214ccdf85d944a9baa66aa70c9997f0; express:sess=eyJ0b2tlbiI6ImEwOWUwY2VhZDE0YTkyYmYxNzhmOWQyYTQ2NTM2MDM5OGVlZDcxNTg3YjY2Y2MyNzU4ZDZmMWJjNGFkODVmODkiLCJzYWx0IjoiMTU5NDE1ODMzNDE3OCJ9; ghost-members-ssr=toddbirchard+jfphhudgshjdiwhekjb@gmail.com; ghost-members-ssr.sig=tM3JLCh6p3msMkt_hAiQmOZl3hg; __stripe_sid=3beca9ec-39cf-42b2-9a75-b3d5fa008faf"
	var headers http.Header
	headers.Add("cookie", cookie)
	if err != nil {
		log.Fatal(err)
	}
	req := &http.Request{URL: endpoint, Header: headers}
	fmt.Println(req)
}

/* func TestGetUserSession(t *testing.T) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request
	req := &http.Request{URL: t.endpoint, Header: headers}
	res, reqError := client.Do(req)
	if reqError != nil {
		log.Fatal(reqError)
	}
	if res.StatusCode != 200 {
		log.Fatal("status code error: %i", res.StatusCode)
	}
	fmt.Println("Hello")
	fmt.Println(res)

	// We want our status to be 200
	if res.StatusCode != http.StatusOK {
		t.Log("Status should be ok, got %i", res.StatusCode)
	}
}*/
