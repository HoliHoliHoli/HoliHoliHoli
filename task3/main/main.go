package main

import (
	"fmt"
	"holi/task3/websites"
	"holi/task3/socialmedia"
	"holi/task3/exporter"

)

func main() {
	fb := new(websites.Facebook)
	twtr := new(websites.Twitter)
	lnkdin := new(websites.Linkedin)
	err := exporter.Export(twtr, []string{"task3/twtrdata.txt", "task3/twtrdata.json", "task3/twtrdata.xml"})
	err = exporter.Export(fb, []string{"task3/fbdata.txt", "task3/fbdata.json", "task3/fbdata.xml"})
	err = exporter.Export(lnkdin, []string{"task3/lnkdin.txt", "task3/lnkdin.json", "task3/lnkdin.xml"})


	if err != nil {
		panic(err)
	}


	ScrollFeeds(fb, twtr, lnkdin)
}

// ScrollFeeds prints all social media feeds
func ScrollFeeds(platforms ...socialmedia.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}