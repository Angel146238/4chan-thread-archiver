// Package fourchan provides an API client for 4chan.org

// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package fourchan

const (
	// API_VERSION is the internal api revision number
	API_VERSION = "0.0.0"

	// HTTP_CLIENT_USER_AGENT is the user agent we use for http requests
	HTTP_CLIENT_USER_AGENT = "4chan-thread-archiver"

	// API_THREAD_URL provides a url to reach a threads info url given:
	// {board} - The name of the board
	// {threadnumber} - The number of the thread
	// Supports GET, HEAD, OPTIONS
	API_THREAD_URL = "https://a.4cdn.org/{board}/thread/{threadnumber}.json"

	// API_BOARD_PAGE_URL provides a url to reach a boards paginated thread index given:
	// {pagenumber} - The page of results, (1 is the main index)
	// Supports GET, HEAD, OPTIONS
	API_BOARD_PAGE_URL = "https://a.4cdn.org/{board}/{pagenumber}.json"

	// API_BOARD_CATALOG_URL provides a url to reach a boards catalog given:
	// {board} - The name of the board
	// Supports GET, HEAD, OPTIONS
	API_BOARD_CATALOG_URL = "https://a.4cdn.org/{board}/catalog.json"

	// API_BOARD_ARCHIVE_URL provides a url to reach a boards archive given:
	// {board} - The name of the board
	// Supports GET, HEAD, OPTIONS
	API_BOARD_ARCHIVE_URL = "https://a.4cdn.org/{board}/archive.json"

	// API_BOARD_THREAD_LIST_URL provides a url to list a boards threads given:
	// {board} - The name of the board
	// Supports GET, HEAD, OPTIONS
	API_BOARD_THREAD_LIST_URL = "https://a.4cdn.org/{board}/threads.json"

	// API_BOARD_LIST_URL provides a url to list all boards
	// Supports GET, HEAD, OPTIONS
	API_BOARD_LIST_URL = "https://a.4cdn.org/boards.json"

	// API_IMAGE_URL provides a url to an image from a post given:
	// {board} - The name of the board
	// {tim} - The renamed filename
	// {ext} - The file extension
	API_IMAGE_URL = "https://i.4cdn.org/{board}/{tim}{ext}"

	// API_IMAGE_THUMBNAIL_URL provides a url to an image thumbnail from a post given:
	// {board} - The name of the board
	// {tim} - The renamed filename
	API_IMAGE_THUMBNAIL_URL = "https://t.4cdn.org/{board}/{tim}s.jpg"

	// FOURCHAN_TIME_FORMAT is the time format that appears to be used across 4chan
	// In the API docs, it is given as:
	// `MM\/DD\/YY(Day)HH:MM (:SS on some boards), EST/EDT timezone`
	FOURCHAN_TIME_FORMAT = "01\\/02\\/06(Mon)15:04:05"
)
