package eduProject

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int    `json:"id"`
	UserId string `json:"userid"`
	ListId int    `json:"listid"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ListId int `json:"listid"`
	ItemId int `json:"itemid"`
}
