package post_like

import (
	"errors"
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeletePostLikeRequestBody struct {
	PostLikeID uuid.UUID `json:"PostLikeID"`
}

func (h handler) DeletePostLike(c *gin.Context) {
	body := DeletePostLikeRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, ok := c.Get("UserID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is missing from context"})
		return
	}

	var postLike models.PostLike

	if result := h.DB.Find(&postLike, body.PostLikeID); result.Error != nil || postLike.ID == uuid.Nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if postLike.UserID != userID {
		c.AbortWithError(http.StatusUnauthorized, errors.New(""))
		return
	}

	if result := h.DB.Delete(&postLike); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, &postLike)
}
