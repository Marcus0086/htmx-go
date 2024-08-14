package errors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHTTPError(c *gin.Context, err error) {
	log.Println("Error:", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
}
