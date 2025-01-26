package dtos

type CreateProductDto struct {
	PhotoURL      string  `json:"photo_url" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	StockQuantity int     `json:"stock_qnt" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
}

type UpdateProductDto struct {
	PhotoURL      string  `json:"photo_url" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	StockQuantity int     `json:"stock_qnt" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
}
