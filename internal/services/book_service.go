package services

import (
	"github.com/thechervonyashiy/bookstore/internal/dtos"
	"github.com/thechervonyashiy/bookstore/internal/repositories"
)

type BookService interface {
	GetAllBooks() ([]dtos.BookDTO, error)
	CreateBook(title string, author string) (int64, error)
	GetBookByID(id int) (dtos.BookDTO, error)
	DeleteBook(id int64) (int64, error)
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]dtos.BookDTO, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var bookDTOs []dtos.BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, dtos.BookDTO{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
		})
	}
	return bookDTOs, nil
}

func (s *bookService) CreateBook(title string, author string) (int64, error) {
	return s.repo.CreateBook(title, author)
}

func (s *bookService) GetBookByID(id int) (dtos.BookDTO, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return dtos.BookDTO{}, err
	}

	var bookDTOs dtos.BookDTO = dtos.BookDTO{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}

	return bookDTOs, nil
}

func (s *bookService) DeleteBook(id int64) (int64, error) {
	return s.repo.DeleteBook(id)
}
