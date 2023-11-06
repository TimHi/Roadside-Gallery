package data

import (
	"encoding/base64"
	"encoding/csv"
	"log"

	"os"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func GetImageData(file string) string {
	// TODO_THL: env file
	imageFile, err := os.ReadFile("/Users/hiller/dev/gallery/backend/assets/us_roadside_attractions_images/us_roadside_attractions_images/" + file)
	if err != nil {
		log.Fatal("Error reading the image file:", err)
	}

	// Encode the image data as base64
	base64Image := base64.StdEncoding.EncodeToString(imageFile)
	return base64Image
}
