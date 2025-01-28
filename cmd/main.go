package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

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
	var features []model.Feature
	featureCollection := model.FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}

	folders, err := os.ReadDir(inputDirectory)
	if err != nil {
		log.Fatal("could not read input folder content")
	}

	folderBar := progressbar.Default(int64(len(folders)))

	for _, folder := range folders {
		folderBar.Add(1)
		if folder.IsDir() {
			subfolders, err := os.ReadDir(path.Join(inputDirectory, folder.Name()))
			if err != nil {
				log.Fatal("could not read sub folder content")
			}

			for _, subfolder := range subfolders {
				if !subfolder.IsDir() {
					places := readFile(path.Join(inputDirectory, folder.Name(), subfolder.Name()))

					for _, place := range places {
						if place.Population != 0 {
							feature := model.NewFeature("Point", place)
							featureCollection.Features = append(featureCollection.Features, feature)
						}
					}

				}
			}
		}
	}

	file, _ := json.Marshal(featureCollection)

	_ = os.WriteFile("test.json", file, 0644)
}
