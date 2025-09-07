package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Allhuha struct {
	Id      int    `json:"ID"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
	Phone   string `json:"Phone"`
}
