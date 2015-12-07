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

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/HenrySlawniak/4chan-thread-archiver/fourchan"
	"io/ioutil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var threads = []string{
	"po:533882",
}

type ThreadData struct {
	Images []Image
}

type Image struct {
	Name     string
	MD5      string
	Thread   int
	Location string
	Width    int
	Height   int
	Ext      string
	Url      string
}

func main() {
	for {
		DumpThreads()
	}
}

func DumpThreads() {
	for _, str := range threads {
		splat := strings.Split(str, ":")
		board := splat[0]
		threadnum, _ := strconv.Atoi(splat[1])
		thread, err := fourchan.GetThread(board, threadnum)
		if err != nil {
			debug.PrintStack()
			fmt.Println(err.Error())
			continue
		}
		folder := fmt.Sprintf("download/%s/%d/", board, threadnum)
		os.MkdirAll(folder, os.ModeDir)

		threadData := ThreadData{}
		for _, post := range thread.Posts {
			if post.Tim > 0 {
				url := fourchan.FormatImageURL(board, post.Tim, post.Ext)
				fileloc := fmt.Sprintf("%sassets/%d%s", folder, post.Tim, post.Ext)
				realmd5, _ := base64.StdEncoding.DecodeString(post.FileMD5)
				img := Image{
					Name:     fmt.Sprintf("%d%s", post.Tim, post.Ext),
					MD5:      fmt.Sprintf("%x", realmd5),
					Thread:   threadnum,
					Location: fileloc,
					Width:    post.Width,
					Height:   post.Height,
					Ext:      post.Ext,
					Url:      url,
				}
				threadData.Images = append(threadData.Images, img)
				fourchan.DownloadFile(fileloc, url, "*", fourchan.HTTP_CLIENT_USER_AGENT)
			}
		}
		j, _ := json.MarshalIndent(threadData, "", "    ")
		ioutil.WriteFile(fmt.Sprintf("%s%d.json", folder, threadnum), j, 0777)
		time.Sleep(1 * time.Second)
	}
}
