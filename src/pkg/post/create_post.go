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
    var user models.User
    
    if  result := h.DB.Find(&user, body.UserID); result.Error != nil || user.ID <= 0  {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }


    post.User = user
    post.Content = body.Content

    if result := h.DB.Create(&post); result.Error != nil {
        c.AbortWithError(http.StatusInternalServerError, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &post)
}