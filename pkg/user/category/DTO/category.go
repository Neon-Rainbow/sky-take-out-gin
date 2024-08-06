package DTO

import model "sky-take-out-gin/model/sql"

type CategoryRequestDTO struct {
	Type int `form:"type" binding:"required"`
}

type CategoryResponseDTO struct {
	Categories []model.Category `json:"categories"`
}
