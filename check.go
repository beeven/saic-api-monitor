package main

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)


func checkLogin() error {
	q, _ := query.Values(config.Login)
	loginUrl := config.API.BaseUrl + config.API.LoginPath + "?" + q.Encode()
	client := &http.Client{
		Timeout:       time.Duration(5 * time.Second),
	}
	_, err := client.Get(loginUrl)
	if err != nil {
		return err
	}
	return nil
}
