// Package fourchan provides an API client for 4chan.org

// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package fourchan

import (
	"fmt"
	"strings"
)

// FormatThreadURL returns an api url for the thread
func FormatThreadURL(board string, number int) string {
	s := strings.Replace(API_THREAD_URL, "{board}", board, 1)
	s = strings.Replace(s, "{threadnumber}", fmt.Sprintf("%d", number), 1)
	return s
}

// FormatThreadURL returns an api url for the board index page
func FormatBoardPageURL(board string, page int) string {
	s := strings.Replace(API_BOARD_PAGE_URL, "{board}", board, 1)
	s = strings.Replace(s, "{pagenumber}", fmt.Sprintf("%d", page), 1)
	return s
}

// FormatBoardCatalogURL returns the api url for the board catalog
func FormatBoardCatalogURL(board string) string {
	s := strings.Replace(API_BOARD_CATALOG_URL, "{board}", board, 1)
	return s
}

// FormatBoardCatalogURL returns the api url for the board archive
func FormatBoardArchiveURL(board string) string {
	s := strings.Replace(API_BOARD_ARCHIVE_URL, "{board}", board, 1)
	return s
}

// FormatBoardThreadListURL returns the api url for the board's thread list
func FormatBoardThreadListURL(board string) string {
	s := strings.Replace(API_BOARD_THREAD_LIST_URL, "{board}", board, 1)
	return s
}

// FormatImageURL returns the url for the image on the board
func FormatImageURL(board string, tim int, ext string) string {
	s := strings.Replace(API_IMAGE_URL, "{board}", board, 1)
	s = strings.Replace(s, "{tim}", fmt.Sprintf("%d", tim), 1)
	s = strings.Replace(s, "{ext}", ext, 1)
	return s
}

// FormatImageThumbnailURL returns the url for the image's thumbnail
func FormatImageThumbnailURL(board string, tim int) string {
	s := strings.Replace(API_IMAGE_THUMBNAIL_URL, "{board}", board, 1)
	s = strings.Replace(s, "{tim}", fmt.Sprintf("%d", tim), 1)
	return s
}

// FormatBoardListURL returns the api url for the board list
func FormatBoardListURL() string {
	return API_BOARD_LIST_URL
}
