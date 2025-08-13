package response

import "github.com/gin-gonic/gin"

func OK(c *gin.Context, data any)              { c.JSON(200, gin.H{"data": data}) }
func Created(c *gin.Context, data any)         { c.JSON(201, gin.H{"data": data}) }
func BadRequest(c *gin.Context, msg string)    { c.JSON(400, gin.H{"error": msg}) }
func Unauthorized(c *gin.Context, msg string)  { c.JSON(401, gin.H{"error": msg}) }
func NotFound(c *gin.Context, msg string)      { c.JSON(404, gin.H{"error": msg}) }
func Internal(c *gin.Context, msg string)      { c.JSON(500, gin.H{"error": msg}) }
