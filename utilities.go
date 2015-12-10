// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"github.com/HenrySlawniak/4chan-thread-archiver/fourchan"
)

// GetVersionString returns the version string
func GetVersionString() string {
	return VERSION + " api-" + fourchan.API_VERSION + "-git-" + GitDescribe()
}

// PrintLicense prints the license information to stdout
func PrintLicense() {
	fmt.Println("")
	fmt.Println("Starting 4chan Thread Archiver version: " + GetVersionString())
	fmt.Println("Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>")
	fmt.Println("SPDX-License-Identifier: MIT")
	fmt.Println("")
}
