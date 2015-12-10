// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/HenrySlawniak/4chan-thread-archiver/fourchan"
	"html/template"
	"io/ioutil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

const VERSION = "0.0.0"

var threadTemplate *template.Template
var box *rice.Box

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
	PrintLicense()
	SetupTemplates()

	for {
		DumpThreads()
	}
}

func SetupTemplates() {
	var err error
	box, err = rice.FindBox("templates")
	if err != nil {
		panic(err)
	}

	threadTemplate, err = template.New("thread").Parse(box.MustString("thread.html")) //.ParseFiles("./templates/thread.html")
	if err != nil {
		panic(err)
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
		threadFileloc := fmt.Sprintf("%s/%d.html", folder, threadnum)
		threadFile, err := os.OpenFile(threadFileloc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			debug.PrintStack()
			fmt.Println(err.Error())
			continue
		}
		defer threadFile.Close()
		threadTemplate.Execute(threadFile, map[string]interface{}{"posts": thread.Posts})
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
