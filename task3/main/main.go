package main

import (
	"fmt"
	"holi/task3/websites"
	"holi/task3/socialmedia"
	"os"
	"errors"
)

func main() {
	fb := new(websites.Facebook)
	twtr := new(websites.Twitter)
	lnkdin := new(websites.Linkedin)
	err := export(twtr, "twtrdata.txt")
	err = export(fb, "fbdata.txt")
	err = export(lnkdin, "lnkdin.txt")

	if err != nil {
		panic(err)
	}


	ScrollFeeds(fb, twtr, lnkdin)
}

// ScrollFeeds prints all social media feeds
func ScrollFeeds(platforms ...socialmedia.SocialMedia) {
	for _, sm := range platforms {
		for a, fd := range sm.Feed() {
			fmt.Println(a, fd)
		}
		fmt.Println("=================================")
	}
}

func export(u socialmedia.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

//    new new

type xmlStruct struct {

}
