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
	err := exporter.Export(twtr, []string{"twtrdata.txt", "twtrdata.json"})
	err = exporter.Export(fb, []string{"fbdata.txt", "fbdata.json"})
	err = exporter.Export(lnkdin, []string{"lnkdin.txt", "lnkdin.json"})


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