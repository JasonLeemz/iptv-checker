package dto

type LiveSourceDTO struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
	Data   string `json:"data"`
	Ctime  string `json:"ctime"`
	Utime  string `json:"utime"`
}
