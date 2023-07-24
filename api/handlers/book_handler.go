package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/your-username/bookstore_microservice/api/models"
)

var mutex = sync.RWMutex{}
var lastID = 0
var books = make(map[int]models.Book)

// GetAllBooks
// @Summary Get a list of all books
// @Description Returns a list of all books in the bookstore.
// @Produce json
// @Success 200 {array} Book
// @Router /books [get]
func GetAllBooks(c echo.Context) error {
	mutex.RLock()
	defer mutex.RUnlock()

	bookList := make([]models.Book, len(books))
	i := 0
	for _, book := range books {
		bookList[i] = book
		i++
	}

	return c.JSON(http.StatusOK, bookList)
}

// GetBook
// @Summary Get details of a book
// @Description Returns details of a specific book.
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} Error
// @Router /books/{id} [get]
func GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	mutex.RLock()
	book, exists := books[id]
	mutex.RUnlock()

	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

// AddBook
// @Summary Add a new book
// @Description Adds a new book with the provided details.
// @Accept json
// @Produce json
// @Param book body Book true "Book details"
// @Success 201 {object} Book
// @Failure 400 {object} Error
// @Router /books [post]
func AddBook(c echo.Context) error {
	var book models.Book
	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if book.Title == "" || book.Author == "" || book.PublishedYear == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Title, Author, and PublishedYear are required")
	}

	mutex.Lock()
	lastID++
	book.ID = lastID
	books[book.ID] = book
	mutex.Unlock()

	return c.JSON(http.StatusCreated, book)
}
