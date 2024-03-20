package request

type MovieRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
	Image       string  `json:"image"`
}

type UpdateMovieRequest struct {
	Id          int     `json:"id" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
	Image       string  `json:"image"`
}

type MovieId struct {
	Id string `uri:"id"`
}
