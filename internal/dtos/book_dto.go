package dtos

type BookDTO struct {
	ID     int    `json:"id"` // Для передачи клиенту
	Title  string `json:"title"`
	Author string `json:"author"`
}
