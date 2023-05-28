package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	// get the contents of the directory
	oldFiles, _ := os.ReadDir("./")

	// rename old files
	for i, file := range oldFiles {
		if err := exec.Command("mv", file.Name(), "v"+strconv.Itoa(i)+".webm").Run(); err != nil {
			log.Fatal(err, "Could not change file name")
		}
	}

	newFiles, _ := os.ReadDir("./")

	// create a file to append content
	f, _ := os.OpenFile("videos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// close file at the end of the program
	defer f.Close()

	// store file names into videos file
	for _, file := range newFiles {
		if _, err := f.WriteString("file " + file.Name() + "\n"); err != nil {
			log.Fatal(err)
		}
	}

	// create a script file
	sf, _ := os.OpenFile("merge.sh", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if _, err := sf.WriteString(`#!/usr/bin/env bash \n
        ffmpeg -f concat -safe 0 -i videos.txt -c copy output.webm \n`); err != nil {
		log.Fatal(err, "Could not write to file")
	}

	// close script file
	defer sf.Close()

	fmt.Println("Merging videos")

	// Build a spinner for showing progress
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	s.Start()
	if err := exec.Command("bash", "./merge.sh").Run(); err != nil {
		log.Fatalf("ffmpeg: %s", err)
	}

	if err := exec.Command("rm", "videos.txt", "merge.sh").Run(); err != nil {
		log.Fatal(err)
	}
	s.Stop()

	fmt.Print("Do you want to delete individual video files [y/N]: ")

	var input string
	fmt.Scanln(&input)

	if input == "y" {
		s.Start()
		for _, file := range newFiles {
			if err := exec.Command("rm", file.Name()).Run(); err != nil {
				log.Fatal(err, "Could not remove file")
			}
		}
		s.Stop()

		fmt.Println("Deleted individual files successfully")
	} else {
		for i := 0; i < len(oldFiles); i++ {
			if err := exec.Command("mv", newFiles[i].Name(), oldFiles[i].Name()).Run(); err != nil {
				log.Fatal(err, "Could not rename file")
			}
		}
	}

	fmt.Println("Merge Successful")
}
