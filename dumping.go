// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/HenrySlawniak/4chan-thread-archiver/fourchan"
	"io/ioutil"
	"os"
	"runtime/debug"
)

func DumpThread(in ThreadInput) error {
	thread, err := fourchan.GetThread(in.Board, in.Number)
	if err != nil {
		return err
	}

	folder := fmt.Sprintf("download/%s/%d/", in.Board, in.Number)
	os.MkdirAll(folder, os.ModeDir)

	threadData := ThreadData{}
	threadFileloc := fmt.Sprintf("%s/%d.html", folder, in.Number)
	threadFile, err := os.OpenFile(threadFileloc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer threadFile.Close()

	threadTemplate.Execute(threadFile, map[string]interface{}{"posts": thread.Posts})
	for _, post := range thread.Posts {
		if post.Tim > 0 {
			url := fourchan.FormatImageURL(in.Board, post.Tim, post.Ext)
			fileloc := fmt.Sprintf("%sassets/%d%s", folder, post.Tim, post.Ext)
			realmd5, _ := base64.StdEncoding.DecodeString(post.FileMD5)
			img := Image{
				Name:     fmt.Sprintf("%d%s", post.Tim, post.Ext),
				MD5:      fmt.Sprintf("%x", realmd5),
				Thread:   in.Number,
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
	ioutil.WriteFile(fmt.Sprintf("%s%d.json", folder, in.Number), j, 0777)
	return nil
}

func DumpThreads() {
	for _, in := range threads {
		err := DumpThread(in)
		if err != nil {
			debug.PrintStack()
			fmt.Println(err.Error())
		}
	}
}
