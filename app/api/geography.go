package api

import (
	geographyHandler "event-reporting/app/handler/geography"
	middleware "event-reporting/app/helpers/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

type geographyGroup struct {
	RouterGroup *gin.RouterGroup
}

func (r *geographyGroup) Init() {
	defer func() {
		fmt.Println("Geography API has been initialized")
	}()
	r.RouterGroup.GET("/geography/countries/", middleware.JWT(), geographyHandler.SearchCountries)
	r.RouterGroup.GET("/geography/states/", middleware.JWT(), geographyHandler.SearchStates)
	r.RouterGroup.GET("/geography/districts/", middleware.JWT(), geographyHandler.SearchDistricts)
	r.RouterGroup.GET("/geography/cities/", middleware.JWT(), geographyHandler.SearchCities)
}
