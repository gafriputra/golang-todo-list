package activity

type GetActivityDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateActivityInput struct {
	Email string `json:"email" binding:"required"`
	Title string `json:"title" binding:"required"`
}
