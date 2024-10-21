package todo

type TodoList struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

type UserList struct {
	ID     int64 `json:"id"`
	UserId int64 `json:"user_id"`
	TypeId int64 `json:"type"`
}

type TodoItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	ID     int64 `json:"id"`
	ListId int64 `json:"list_id"`
	ItemId int64 `json:"item_id"`
}
