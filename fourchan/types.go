// Package fourchan provides an API client for 4chan.org

// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package fourchan

import "html/template"

// Thread represents a 4chan thread, a slice of posts.
type Thread struct {
	Posts []Post
}

// Post represents a 4chan post, chceck api.md for more info on it's structure.
type Post struct {
	Number         int              `json:"no"`              // Post number
	Resto          int              `json:"resto"`           // Reply to
	Sticky         int              `json:"sticky"`          // Stickied thread?
	Closed         int              `json:"closed"`          // Closed thread?
	Archived       int              `json:"archived"`        // Archived thread?
	Now            string           `json:"now"`             // Date and time
	Time           int              `json:"time"`            // UNIX timestamp
	Name           string           `json:"name"`            // Name
	Trip           string           `json:"trip"`            // Tripcode
	Id             string           `json:"id"`              // ID
	Capcode        string           `json:"capcode"`         // Capcode
	Country        string           `json:"country"`         // Country Code
	CountryName    string           `json:"country_name"`    // Country Name
	Sub            string           `json:"sub"`             // Subject
	Com            template.HTML    `json:"com"`             // Comment (escaped HTML)
	Tim            int              `json:"tim"`             // Renamed filename
	Filename       string           `json:"filename"`        // Original filename
	Ext            string           `json:"ext"`             // File extenstion
	Fsize          int              `json:"fsize"`           // File size
	FileMD5        string           `json:"md5"`             // File MD5
	Width          int              `json:"w"`               // Image width
	Height         int              `json:"h"`               // Image height
	ThumbWidth     int              `json:"tn_w"`            // Thumbnail width
	ThumbHeight    int              `json:"tn_h"`            // Thumbnail height
	FileDeleted    int              `json:"filedeleted"`     // File deleted?
	Spoiler        int              `json:"spoiler"`         // Spoiler image?
	CustomSpoiler  int              `json:"custom_spoiler"`  // Custom spoilers?
	OmmitedPosts   int              `json:"omitted_posts"`   // # replies ommited
	OmmitedImages  int              `json:"omitted_images"`  // # image replies ommited
	Replies        int              `json:"replies"`         // # replies total
	Images         int              `json:"images"`          // # images total
	BumpLimit      int              `json:"bumplimit"`       // Bump limit met?
	ImageLimit     int              `json:"imagelimit"`      // image limit met?
	CapcodeReplies map[string][]int `json:"capcode_replies"` // Capcode user replies?
	LastModified   int              `json:"last_modified"`   // Time when last modified
	Tag            string           `json:"tag"`             // Thread tag
	SemanticUrl    string           `json:"semantic_url"`    // Thread URL slug
}

func (p Post) IsImage() bool {
	return p.Ext == ".png" || p.Ext == ".jpg" || p.Ext == ".gif"
}

func (p Post) IsVideo() bool {
	return p.Ext == ".webm"
}
