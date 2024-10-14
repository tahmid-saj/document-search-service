package routes

import (
	"document-search-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func searchDocuments(context *gin.Context) {
	searchQuery := context.Param("query")

	var searchQueryInput models.SearchQueryInput
	err := context.ShouldBindJSON(&searchQueryInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body"})
		return
	}

	res, err := models.SearchDocuments(searchQuery, searchQueryInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch documents"})
		return
	}

	context.JSON(http.StatusOK, res)
}