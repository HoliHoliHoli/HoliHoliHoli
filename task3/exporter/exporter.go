package exporter

import (
	"fmt"
	"holi/task3/socialmedia"
	"encoding/json"
	"encoding/xml"
	"os"
	"errors")

//Export stuff
func Export(u socialmedia.SocialMedia, filenames []string) error {
	for _, filename := range (filenames){
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return errors.New("an error occured opening the file: " + err.Error())
		}

		// To check the file extension
		switch filename[len(filename)-3:] {
		case "txt":
			for _, data := range u.Feed() {
				n, err := f.Write([]byte(data + "\n"))
				if err != nil {
					return errors.New("an error occured writing to file: " + err.Error())
				}
				fmt.Printf("wrote %d bytes\n", n)
			}

		case "son":
			for index, data := range u.Feed() {
				line := map[int]string{index:data}
				jsonline, _ := json.Marshal(line)
				n, err := f.Write([]byte(string(jsonline) + "\n"))
				if err != nil {
					return errors.New("an error occured writing to file: " + err.Error())
				}
				fmt.Printf("wrote %d bytes\n", n)
			}

		case "xml":
			for _, fd := range u.Feed(){
				xmlline, _ := xml.MarshalIndent(fd, "", "")
				n, err := f.Write([]byte(string(xmlline) + "\n"))
				if err != nil {
					return errors.New("an error occured writing to file: " + err.Error())
				}
				fmt.Printf("wrote %d bytes\n", n)
			}
			
		}

	}
	return nil
}

