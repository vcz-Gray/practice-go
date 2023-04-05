package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"test/go/config"
	"test/go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	body := c.Request.Body
	b, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"userInfo": "Wrong approach",
		})
	}

	var u models.User
	json.Unmarshal([]byte(b), &u)
	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(password)

	config.GetDB().Create(&u)
	c.JSON(http.StatusOK, gin.H{
		"userInfo": u,
	})
}
