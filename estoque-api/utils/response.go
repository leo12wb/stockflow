package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PagedResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	LastPage int         `json:"last_page"`
}

type Pagination struct {
	Page    int
	PerPage int
	Offset  int
}

func ParsePagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	return Pagination{
		Page:    page,
		PerPage: perPage,
		Offset:  (page - 1) * perPage,
	}
}

func PagedSuccessResponse(c *gin.Context, data interface{}, total int64, p Pagination) {
	lastPage := int(total) / p.PerPage
	if int(total)%p.PerPage != 0 {
		lastPage++
	}
	c.JSON(http.StatusOK, PagedResponse{
		Success:  true,
		Data:     data,
		Total:    total,
		Page:     p.Page,
		PerPage:  p.PerPage,
		LastPage: lastPage,
	})
}

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   message,
	})
}

func ValidationErrorResponse(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Error:   err,
	})
}
