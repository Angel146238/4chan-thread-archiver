// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package main

import "os/exec"

func GitDescribe() string {
	out, err := exec.Command("git", "describe", "--always", "--dirty=*").Output()
	if err != nil {
		return "?"
	}
	return string(out)[:len(out)-1]
}
