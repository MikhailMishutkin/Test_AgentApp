package model

type MetroList struct {
	CityId   string  `json:"id"`
	Lines    []Lines `json:"lines"`
	CityName string  `json:"name"`
	URLCity  string  `json:"url"`
}

type Lines struct {
	Color    string     `json:"hex_color"`
	LineID   string     `json:"id"`
	LName    string     `json:"name"`
	Stations []Stations `json:"stations"`
}

type Stations struct {
	StID   string  `json:"id"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	StName string  `json:"name"`
	Order  int     `json:"order"`
}
