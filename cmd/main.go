package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pbusenius/ndjson-converter/model"
	"github.com/schollz/progressbar/v3"
)

const inputDirectory = "data"

func readFile(path string) []model.Place {
	var places []model.Place

	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var place model.Place
		err = json.Unmarshal([]byte(fileScanner.Text()), &place)
		if err != nil {
			log.Fatalf("could not parse json string %v", err)
		}

		places = append(places, place)
	}

	return places
}

func main() {
	var citiyFiles []string
	var features []model.Feature
	featureCollection := model.FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}

	folders, err := os.ReadDir(inputDirectory)
	if err != nil {
		log.Fatal("could not read input folder content")
	}

	err = filepath.Walk(inputDirectory, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.Contains(path, ".ndjson") {
			citiyFiles = append(citiyFiles, path)
		}

		return nil
	})
	if err != nil {
		log.Fatal("could not walk directory")
	}

	cityBar := progressbar.Default(int64(len(folders)))

	for _, cityFile := range citiyFiles {
		cityBar.Add(1)
		places := readFile(cityFile)

		for _, place := range places {
			if place.Population != 0 {
				feature := model.NewFeature("Point", place)
				featureCollection.Features = append(featureCollection.Features, feature)
			}
		}
	}

	file, _ := json.Marshal(featureCollection)

	_ = os.WriteFile("test.json", file, 0644)
}
