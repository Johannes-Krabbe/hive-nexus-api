package post

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
)

type CreatePostRequestBody struct {
    UserID uint `json:"UserID"`
    Content string `json:"Content"`
}

func (h handler) CreatePost(c *gin.Context) {
    body := CreatePostRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var post models.Post

    post.User = models.User{ID: body.UserID}
    post.Content = body.Content

    if result := h.DB.Create(&post); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &post)
}