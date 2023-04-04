package main

import (
	"test/go/users"

	"github.com/gin-gonic/gin"
)

var latest map[string]int = make(map[string]int)

// var article map[int]Article = make(map[int]Article)

func main() {
	r := gin.Default()
	crudUser(r)
	// crudArticle(r)
	r.Run()
}

func crudUser(g *gin.Engine) {
	g.GET("user/:id", users.GetUserById)
	g.POST("user/", users.PostUserById)
	g.PUT("user/:id", users.PutUserById)
	g.DELETE("user/:id", users.DeleteUserById)
}
