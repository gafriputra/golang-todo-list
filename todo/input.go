package todo

type GetTodoDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTodoInput struct {
	ActivityGroupID int    `json:"activity_group_id" binding:"required"`
	Title           string `json:"title" binding:"required"`
}
