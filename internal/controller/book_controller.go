package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/actions"
	"github.com/savioafs/book-market/internal/common"
	"github.com/savioafs/book-market/internal/dto"
	"github.com/savioafs/book-market/internal/usecase"
	"net/http"
	"strconv"
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
	if errors.Is(err, common.ErrBookAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
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

func (ct *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id cannot empty",
		})

		return
	}

	book, err := ct.useCase.GetBookByID(id)
	if errors.Is(err, common.ErrBookNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get book by id",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (ct *BookController) GetBooksByCategory(c *gin.Context) {
	category := c.Param("category")

	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "category cannot empty",
		})

		return
	}

	books, err := ct.useCase.GetBookByCategory(category)
	if errors.Is(err, common.ErrCategoryNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "cannot get books by category",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get books by category",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, books)
}

func (ct *BookController) GetBooksByPublishedYear(c *gin.Context) {
	publishedYear := c.Param("published_year")

	if publishedYear == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "published year cannot empty",
		})

		return
	}

	publishedYearConv, err := strconv.Atoi(publishedYear)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot convert published year",
			"error":   err.Error(),
		})
		return
	}

	books, err := ct.useCase.GetBooksByPublishedYear(publishedYearConv)
	if errors.Is(err, common.ErrPublishedYearNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "cannot get books by category",
			"error":   err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get books by category",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (ct *BookController) GetBooksByAuthor(c *gin.Context) {
	author := c.Param("author")

	if author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "category cannot empty",
		})

		return
	}

	books, err := ct.useCase.GetBooksByAuthor(author)
	if errors.Is(err, common.ErrAuthorNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "cannot get books by author",
			"error":   err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get books by category",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (ct *BookController) UpdateBook(c *gin.Context) {
	var book dto.BookInputDTO

	id := c.Param("id")

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid book input",
			"error":   err.Error(),
		})
		return
	}

	err = ct.useCase.UpdateBook(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot update book",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated with success",
	})
}

func (ct *BookController) UpdateStockBookRenew(c *gin.Context) {
	id := c.Param("id")
	quantity := c.Query("quantity")

	quantityConvert, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error to convert quantity to int",
			"error":   err.Error(),
		})
		return
	}

	err = ct.useCase.UpdateStockBook(id, quantityConvert, actions.BookRenew)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot update stock renew",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update stock renew with success",
	})
}

func (ct *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := ct.useCase.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "does not delete book",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete with success",
	})

}
