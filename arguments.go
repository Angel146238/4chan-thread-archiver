// Copyright (c) 2015 Henry Slawniak <henry@slawniak.com>
// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ThreadInput struct {
	Board  string
	Number int
}

var threads = []ThreadInput{}

func LoadThreadsFromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		splat := strings.Split(i, ":")
		if len(splat) < 2 {
			fmt.Println(i + " is not a valid 4chan thread input")
			continue
		}
		num, err := strconv.Atoi(splat[1])
		if err != nil {
			fmt.Println(i + " is not a valid 4chan thread input")
			continue
		}
		input := ThreadInput{
			Board:  splat[0],
			Number: num,
		}
		threads = append(threads, input)
	}

}

func LoadThreadsFromArgs(args []string) {
	for _, i := range args {
		splat := strings.Split(i, ":")
		if len(splat) < 2 {
			fmt.Println(i + " is not a valid 4chan thread input")
			continue
		}
		num, err := strconv.Atoi(splat[1])
		if err != nil {
			fmt.Println(i + " is not a valid 4chan thread input")
			continue
		}
		input := ThreadInput{
			Board:  splat[0],
			Number: num,
		}
		threads = append(threads, input)
	}
}
