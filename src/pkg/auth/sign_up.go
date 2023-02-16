package auth

import (
	"fmt"
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
)

type SignUpRequestBody struct {
    Username        string `json:"username"`
    Email           string `json:"email"`
    Password        string `json:"password"`
}

func (h handler) SignUp(c *gin.Context) {
    body := SignUpRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.User

	//checking if user with username or email already exists
	// using .DB.Limit(1).Find here instead of .First to prevent error messages
    if  h.DB.Limit(1).Find(&user,"username = ?", body.Username); user.ID > 0  {
		fmt.Println(user)
        c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error": true, "message":"Username already exists"})
        return
    }

    if  h.DB.Limit(1).Find(&user,"email = ?", body.Email); user.ID > 0  {
        c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error": true, "message":"Email already exists"})
        return
    }

	// creating the user
    user.Username = body.Username
    user.Email = body.Email

	user.Salt = getSalt()

	pw, err := hash(body.Password, user.Salt)
	if(err != nil){
        c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error": true, "message":"Something went wrong"})
        return
	}

    user.Password = pw

	fmt.Println("password", user.Password)
	fmt.Println("salt", user.Salt)


    if result := h.DB.Create(&user); result.Error != nil {
        c.AbortWithError(http.StatusInternalServerError, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &user)
}