// Package fourchan provides an API client for 4chan.org

// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
