package main

import (
	"test/go/controller"
	"test/go/routes"

	"github.com/gin-gonic/gin"
)

var latest map[string]int = make(map[string]int)

func main() {
	r := gin.Default()
	crudUser(r)
	crudArticle(r)
	r.Run()
}

func crudUser(g *gin.Engine) {
	g.GET("user/:id", routes.GetUserById)
	g.POST("user/register", controller.RegisterUser)
	g.PUT("user/:id", routes.UpdaeUserById)
	g.DELETE("user/:id", routes.DeleteUserById)
}

func crudArticle(g *gin.Engine) {
	g.GET("article/:id", routes.GetPostById)
	g.POST("article/post", routes.CreateNewPost)
	g.PUT("article/put/:id", routes.UpdatePostById)
	g.DELETE("article/:id", routes.DeletePostById)
}
