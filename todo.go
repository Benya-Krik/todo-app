package eduProject

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int    `json:"id"`
	UserId string `json:"user_id"`
	ListId int    `json:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ListId int `json:"list_id"`
	ItemId int `json:"item_id"`
}

type UpdateListItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListItem) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("Update structure has no values")
	}
	return nil
}
