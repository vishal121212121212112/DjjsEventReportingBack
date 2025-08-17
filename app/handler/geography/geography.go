package geographyHandler

import (
	"event-reporting/app/dtos"
	geographyServiceHandler "event-reporting/app/services/geography"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var geographyService geographyServiceHandler.GeographyService

func SearchCountries(c *gin.Context) {
	var q dtos.CommonQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}

	res, err := geographyService.SearchCountries(q.Q, q.Limit, q.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func SearchStates(c *gin.Context) {
	var q dtos.StatesQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}
	// Optional: early UUID validation for nicer 400s (service also validates)
	if q.CountryID != "" && !isUUID(q.CountryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "country_id must be a valid UUID"})
		return
	}

	res, err := geographyService.SearchStates(q.CountryID, q.Q, q.Limit, q.Offset)
	if err != nil {
		status := http.StatusInternalServerError
		if isClientInputError(err) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func SearchDistricts(c *gin.Context) {
	var q dtos.DistrictsQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}
	if q.StateID != "" && !isUUID(q.StateID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "state_id must be a valid UUID"})
		return
	}

	res, err := geographyService.SearchDistricts(q.StateID, q.Q, q.Limit, q.Offset)
	if err != nil {
		status := http.StatusInternalServerError
		if isClientInputError(err) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func SearchCities(c *gin.Context) {
	var q dtos.CitiesQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}
	if q.DistrictID != "" && !isUUID(q.DistrictID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "district_id must be a valid UUID"})
		return
	}

	res, err := geographyService.SearchCities(q.DistrictID, q.Q, q.Limit, q.Offset)
	if err != nil {
		status := http.StatusInternalServerError
		if isClientInputError(err) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func isUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

// If your service returns errors like "invalid country_id: ...", map them to 400
func isClientInputError(err error) bool {
	msg := err.Error()
	return containsAny(msg, []string{
		"invalid country_id",
		"invalid state_id",
		"invalid district_id",
	})
}

func containsAny(s string, needles []string) bool {
	for _, n := range needles {
		if n != "" && containsFold(s, n) {
			return true
		}
	}
	return false
}

func containsFold(s, sub string) bool {
	// case-insensitive contains without importing extra deps
	ls, lsub := []rune(s), []rune(sub)
	n, m := len(ls), len(lsub)
	if m == 0 || m > n {
		return false
	}
	// naive scan (strings.ContainsFold was added later; keep it portable)
	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			r1 := ls[i+j]
			r2 := lsub[j]
			// normalize ASCII only; if you need full unicode casefolding, use strings.EqualFold on slices
			if toLowerASCII(r1) != toLowerASCII(r2) {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func toLowerASCII(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
