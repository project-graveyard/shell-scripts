package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	// get the contents of the directory
	oldFiles, _ := ioutil.ReadDir("./")

	// rename old files
	for i, file := range oldFiles {
		exec.Command("mv", file.Name(), "v"+strconv.Itoa(i)+".webm").Run()
	}

	newFiles, _ := ioutil.ReadDir("./")

	// create a file to append content
	f, _ := os.OpenFile("videos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// close file at the end of the program
	defer f.Close()

	// store file names into videos file
	for _, file := range newFiles {
		f.WriteString("file " + file.Name() + "\n")
	}

	// create a script file
	sf, _ := os.OpenFile("merge.sh", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	sf.WriteString("#!/usr/bin/env bash \n")
	sf.WriteString("ffmpeg -f concat -safe 0 -i videos.txt -c copy output.webm \n")

	// close script file
	defer sf.Close()

	fmt.Println("Merging videos")

	// Build a spinner for showing progress
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	s.Start()
	if err := exec.Command("bash", "./merge.sh").Run(); err != nil {
		log.Fatalf("ffmpeg: %s", err)
	}
	exec.Command("rm", "videos.txt", "merge.sh").Run()
	s.Stop()

	fmt.Print("Do you want to delete individual video files [y/N]: ")

	var input string
	fmt.Scanln(&input)

	if input == "y" {
		s.Start()
		for _, file := range newFiles {
			exec.Command("rm", file.Name()).Run()
		}
		s.Stop()

		fmt.Println("Deleted individual files successfully")
	} else {
		for i := 0; i < len(oldFiles); i++ {
			exec.Command("mv", newFiles[i].Name(), oldFiles[i].Name()).Run()
		}
	}

	fmt.Println("Merge Successful")
}
