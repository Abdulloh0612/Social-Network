package models

type Post struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	Views    int64  `json:"views"`
	Category string `json:"category"`
	OwnerId  string `json:"owner_id"`
}

type PostCreate struct {
	Content  string `json:"content"`
	Title    string `json:"title"`
	Views    int64  `json:"views"`
	Category string `json:"category"`
	OwnerId  string `json:"owner_id"`
}

type PostReq struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	Views     int64  `json:"views"`
	Category  string `json:"category"`
	OwnerId   string `json:"owner_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
