package models

type Board struct {
	Cells [][]int `json:"cells"`
}

type Game struct {
	ID    string `json:"id"`
	Board Board  `json:"board"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
