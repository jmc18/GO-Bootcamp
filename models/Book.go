package models

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var ArrayBooks = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 3},
	{ID: "2", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 4},
	{ID: "3", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 5},
	{ID: "4", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 6},
}
