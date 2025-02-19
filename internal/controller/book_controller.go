package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
)

type BookController struct {
	useCase usecase.BookUseCase
}

func NewBookController(useCase usecase.BookUseCase) *BookController {
	return &BookController{useCase: useCase}
}

func (ct *BookController) CreateBook(c *gin.Context) {
	var input dto.BookInputDTO

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	output, err := ct.useCase.CreateBook(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot register book",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (ct *BookController) GetAllBooks(c *gin.Context) {

	books, err := ct.useCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot get all books",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (ct *BookController) GetBookByID(c *gin.Context) {}

func (ct *BookController) GetBooksByCategory(c *gin.Context) {}

func (ct *BookController) GetBooksByPublishedYear(c *gin.Context) {}

func (ct *BookController) GetBooksByAuthor(c *gin.Context) {}

func (ct *BookController) UpdateBook(c *gin.Context) {}

func (ct *BookController) UpdateStockBookSale(c *gin.Context) {}

func (ct *BookController) UpdateStockBookRenew(c *gin.Context) {}

func (ct *BookController) DeleteBook(c *gin.Context) {}

/*

func (bc *BookController) CreateBook(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	book, err := bc.bookUseCase.CreateBook()
}

*/
