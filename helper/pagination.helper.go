package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Pagination struct {
	CurrentPage int
	Limit       int
	Skip        int
	Counts      int
}

func GetPagination(c *gin.Context) *Pagination {
	pgnt := Pagination{CurrentPage: 1, Limit: 3}
	page := c.Query("page")
	limit := c.Query("limit")

	if page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		pgnt.CurrentPage = p
		pgnt.Skip = (pgnt.CurrentPage - 1) * pgnt.Limit

	}
	if limit != "" {
		lm, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		pgnt.Limit = lm
	}
	return &pgnt
}
