package handlers

import (
	"net/http"

	"event-reporting/app/services/user"
	jwtutil "event-reporting/app/utils/jwt"
	"event-reporting/app/utils/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc       services.UserService
	jwtSecret string
}

func NewUserHandler(s services.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{svc: s, jwtSecret: jwtSecret}
}

func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", h.register)
	r.POST("/login", h.login)
}

type registerDTO struct{ Name, Email, Password string }
type loginDTO struct{ Email, Password string }

func (h *UserHandler) register(c *gin.Context) {
	var dto registerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	u, err := h.svc.Register(dto.Name, dto.Email, dto.Password)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Created(c, u)
}

func (h *UserHandler) login(c *gin.Context) {
	var dto loginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	u, err := h.svc.Verify(dto.Email, dto.Password)
	if err != nil {
		response.Unauthorized(c, "invalid credentials")
		return
	}

	tok, _ := jwtutil.Sign(u.ID, u.Email, h.jwtSecret)
	c.JSON(http.StatusOK, gin.H{"token": tok})
}
