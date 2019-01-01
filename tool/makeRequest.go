package download

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func MakeRequest(s string) []byte {
	resp, err := http.Get(s)
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	// log.Println(string(body))
	return []byte(body)
}

func WriteToFile(s string, data []byte) error {
	filePerm := 0644
	filename := "temp_image_dump.png"
	return ioutil.WriteFile(filename, data, os.FileMode(filePerm))
}
