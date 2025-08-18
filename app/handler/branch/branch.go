package branchHandler

import (
	"event-reporting/app/helpers/response"
	"event-reporting/app/models"
	branchServiceHandler "event-reporting/app/services/branch"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var branchService branchServiceHandler.BranchService

func SearchBranches(c *gin.Context) {
	var req models.BranchSearchRequest

	// Parse query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		response.SendBadRequestResponse(c, "Invalid query parameters: "+err.Error())
		return
	}

	// Parse pagination parameters
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			req.Limit = &limit
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			req.Offset = &offset
		}
	}

	// Execute search
	results, err := branchService.SearchBranches(req)
	if err != nil {
		log.Printf("Branch search failed: %v", err)
		response.SendErrorResponse(c, http.StatusInternalServerError, "Failed to search branches")
		return
	}

	response.SendSuccessResponse(c, gin.H{
		"branches": results,
		"total":    len(results),
	})
}
