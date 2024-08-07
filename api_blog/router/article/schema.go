package article

type requestArticle struct {
	Title       string  `json:"title" binding:"required,gte=3"`
	Category    string  `json:"category" binding:"required,gte=3"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Description string  `json:"description" binding:"required,gte=6"`
	Content     string  `json:"content" binding:"required,gte=3"`
}
