// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"os"
)

const VERSION = "0.0.0"

var threadTemplate *template.Template
var box *rice.Box
var file = ""

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

	flag.StringVar(&file, "file", "", "an input file of urls, one per line")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please provide the thread(s) to download in the format {board}:{number}, or a file flag")
	}
	if file != "" {
		LoadThreadsFromFile(file)
	} else {
		LoadThreadsFromArgs(os.Args)
	}

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
