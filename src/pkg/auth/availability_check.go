package auth

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ViewAvailabilityCheckData struct {
	Available bool `json:"available"`
}

func (h handler) AvailabilityCheck(c *gin.Context) {

	// getting params
	username := c.Query("username")
	email := c.Query("email")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if username == "" && email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include Username or Email in query params")
		return
	}

	if username != "" && email != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include just one: Username or Email in query params")
		return
	}

	var user models.User
	if username != "" {
		if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "username = ?", username); user.ID == uuid.Nil {
			c.JSON(http.StatusOK, ViewAvailabilityCheckData{Available: true})
			return
		}
	} else {
		if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "Email = ?", email); user.ID == uuid.Nil {
			c.JSON(http.StatusOK, ViewAvailabilityCheckData{Available: true})
			return
		}
	}

	c.JSON(http.StatusOK, ViewAvailabilityCheckData{Available: false})
}
