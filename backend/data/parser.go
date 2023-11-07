package data

import (
	"encoding/base64"
	"encoding/csv"

	"os"

	"github.com/labstack/gommon/log"
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

func GetImageData(file string) (string, error) {
	assetsPath := os.Getenv("ASSETS")
	imageFile, err := os.ReadFile(assetsPath + file)
	if err != nil {
		log.Error("Error reading the image file:", err)
		return "", err
	}

	// Encode the image data as base64
	base64Image := base64.StdEncoding.EncodeToString(imageFile)
	return base64Image, nil
}
