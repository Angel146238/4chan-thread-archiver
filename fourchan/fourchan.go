// Package fourchan provides an API client for 4chan.org

// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package fourchan

import (
	// "github.com/mvdan/xurls"
	"encoding/json"
)

// GetThread returns a *Thread or an error, given the board anf thread number
func GetThread(board string, number int) (*Thread, error) {
	url := FormatThreadURL(board, number)

	body, err := GetHTMLBody(url)
	if err != nil {
		return nil, err
	}

	thread := Thread{}

	err = json.Unmarshal(body, &thread)
	return &thread, err
}
