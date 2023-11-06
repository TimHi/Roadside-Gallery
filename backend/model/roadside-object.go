package model

type RoadsideObject struct {
	Index      int64  `json:"index"`
	Title      string `json:"title"`
	Year       int64  `json:"year"`
	Image      string `json:"image"`
	Image_link string `json:"image_link"`
	Image_data string `json:"image_data"`
}

func NewRoadsideObject(index int64, title string, year int64, image string, image_link string) *RoadsideObject {
	p := new(RoadsideObject)
	p.Index = index
	p.Title = title
	p.Year = year
	p.Image = image
	p.Image_link = image_link
	return p
}
