package repositories

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/thechervonyashiy/bookstore/internal/models"
)

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	CreateBook(title string, author string) (int64, error)
	GetBookByID(id int) (models.Book, error)
	DeleteBook(id int64) (int64, error)
}

type sqliteBookRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(storagePath string) (*sqliteBookRepository, error) {
	const op = "internal.repositories.NewSQLiteRepository"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS books(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	title TEXT,
	author TEXT,
	d_create TEXT,
	d_update TEXT);`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sqliteBookRepository{db: db}, nil
}

func (s *sqliteBookRepository) GetAllBooks() ([]models.Book, error) {
	const op = "storage.sqlite.GetAllBooks"

	rows, err := s.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.DCreate, &book.DUpdate)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		books = append(books, book)
	}
	return books, nil
}

func (s *sqliteBookRepository) CreateBook(title string, author string) (int64, error) {
	const op = "storage.sqlite.CreateBook"

	stmt, err := s.db.Prepare(`INSERT INTO books(title, author, d_create, d_update)
	VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	currentTime := time.Now().Format(time.RFC3339)

	res, err := stmt.Exec(title, author, currentTime, currentTime)
	if err != nil {
		return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s *sqliteBookRepository) GetBookByID(id int) (models.Book, error) {
	const op = "storage.sqlite.GetBookByID"

	stmt, err := s.db.Prepare("SELECT * FROM books WHERE id = ?")
	if err != nil {
		return models.Book{}, fmt.Errorf("%s: %w", op, err)
	}

	var book models.Book
	err = stmt.QueryRow(id).Scan(&book.ID, &book.Title, &book.Author, &book.DCreate, &book.DUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, fmt.Errorf("%s: no book found with id %s", op, strconv.Itoa(id))
		}
		return models.Book{}, fmt.Errorf("%s: scan result: %w", op, err)
	}

	return book, nil
}

func (s *sqliteBookRepository) DeleteBook(id int64) (int64, error) {
	const op = "storage.sqlite.DeleteBook"

	stmt, err := s.db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	cnt, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return cnt, nil
}
