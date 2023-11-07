package api

import (
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/timhi/gallery-backend/m/v2/data"
	"github.com/timhi/gallery-backend/m/v2/model"
)

var roadsideData = []model.RoadsideObject{}

func LoadData() {
	records := data.ReadCsvFile("assets/us_roadside_attractions.csv")
	for i, record := range records {
		if i > 0 {
			roadsideData = append(roadsideData, *model.NewRoadsideObject(parseInt(record[0]), record[1], parseInt(record[2]), record[3], record[4]))
		}
	}
}

func parseInt(s string) int64 {
	pattern := `\d+`

	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Errorf("Error compiling regex: %s\n", err)
	}

	matches := regex.FindAllString(s, -1)
	if len(matches) == 0 {
		log.Error("No Year")
	}

	r, err := strconv.ParseInt(matches[0], 10, 64)
	if err != nil {
		log.Error(err)
	}
	return r
}

func GetRandomRoadsideObject(c echo.Context) error {
	index := rand.Intn(len(roadsideData))
	object := roadsideData[index]
	img, err := data.GetImageData(object.Image)
	if err != nil {
		return err
	}
	object.Image_data = img
	return c.JSON(http.StatusOK, object)
}
