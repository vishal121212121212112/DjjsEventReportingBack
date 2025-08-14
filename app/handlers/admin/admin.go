package adminHandlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"event-reporting/app/models"
	"event-reporting/app/services/admin" // Import the admin service package
	"event-reporting/app/utils/jwt"  // Import the JWT utility package
	"event-reporting/app/utils/response"
)

// Register Admin Handler
func RegisterAdminHandler(c *gin.Context) {
	var dto struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Email        string `json:"email"`
		ContactNumber string `json:"contact_number"`
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	admin := &models.Admin{
		Username:     dto.Username,
		Password:     dto.Password,
		Email:        dto.Email,
		ContactNumber: dto.ContactNumber,
	}

	// Instantiate the service and call the Register method
	adminService := admin.NewAdminService(nil)  // Pass in the correct repository later
	createdAdmin, err := adminService.RegisterAdmin(admin)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, createdAdmin)
}

// Admin Login Handler
func AdminLoginHandler(c *gin.Context) {
	var dto struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Instantiate the service and call the AdminLogin method
	adminService := admin.NewAdminService(nil) // Pass in the correct repository
	admin, err := adminService.AdminLogin(dto.Username, dto.Password)
	if err != nil {
		response.Unauthorized(c, "invalid credentials")
		return
	}

	// Generate JWT token
	token, _ := jwtutil.Sign(admin.ID.String(), admin.Email, "your-jwt-secret")
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Update Admin Password Handler
func UpdateAdminPasswordHandler(c *gin.Context) {
	var dto struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Instantiate the service and call the UpdateAdminPassword method
	adminService := admin.NewAdminService(nil) // Pass in the correct repository
	err := adminService.UpdateAdminPassword(dto.OldPassword, dto.NewPassword)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
