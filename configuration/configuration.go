package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config struct for spectrum.json
type Spectrum []struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content []struct {
		ID   string `json:"id"`
		Desc string `json:"desc"`
	} `json:"content"`
}

// readConfigJson returns spectrum.json configuration file from static dir
func ReadConfigJson(path string) Spectrum {
	if path == "" {
		path = "static/spectrum.json"
	}
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Spectrum
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
