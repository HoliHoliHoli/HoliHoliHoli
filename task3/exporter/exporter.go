package exporter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"holi/task3/websites"
	"holi/task3/socialmedia"
)


func export(u socialmedia.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	switch filename[len(filename)-3:] {
	case "son":
		
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