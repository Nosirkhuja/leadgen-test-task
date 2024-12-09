package model

type Building struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	City       string `json:"city" binding:"required"`
	Year       int    `json:"year" binding:"required"`
	FloorCount int    `json:"floor_count" binding:"required"`
}
