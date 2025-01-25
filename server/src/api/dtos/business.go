package dtos

type colorSchema struct {
	Primary   string `json:"primary" binding:"required"`
	Secondary string `json:"secondary" binding:"required"`
	Paper     string `json:"paper" binding:"required"`
	Text      string `json:"text" binding:"required"`
}

type CreateBusinessDto struct {
	Name        string      `json:"name" binding:"required"`
	Specialty   string      `json:"specialty" binding:"required"`
	History     string      `json:"history" binding:"required"`
	ColorSchema colorSchema `json:"color_schema" binding:"required"`
}

type UpdateBusinessDto struct {
	Name        string      `json:"name" binding:"required"`
	Specialty   string      `json:"specialty" binding:"required"`
	History     string      `json:"history" binding:"required"`
	ColorSchema colorSchema `json:"color_schema" binding:"required"`
}
