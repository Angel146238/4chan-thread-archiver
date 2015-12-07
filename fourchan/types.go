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
	Com            string           `json:"com"`             // Comment (escaped HTML)
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
