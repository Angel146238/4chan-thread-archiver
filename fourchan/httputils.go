// Package fourchan provides an API client for 4chan.org

// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package fourchan

import (
	"io/ioutil"
	"net/http"
)

// GetHTMLBody returns the body of an HTTP GET request as a []byte, or an error
func GetHTMLBody(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil
	}

	req.Header.Set("User-Agent", HTTP_CLIENT_USER_AGENT)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetHTMLHead returns the response of an HTTP HEAD request or an error
func GetHTMLHead(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, nil
	}

	req.Header.Set("User-Agent", HTTP_CLIENT_USER_AGENT)
	return client.Do(req)
}
