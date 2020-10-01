package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func AboutUs(c *gin.Context) {
	c.HTML(http.StatusOK, "about_us.tmpl", gin.H{})
}
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}
func SubmitLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("******** FROM SubmitLogin ****************")
	fmt.Println("********", email)
	fmt.Println("********", password) // lets print out what we get

	c.Redirect(http.StatusFound, "/")
}

func Custom404(c *gin.Context) {
	c.HTML(404, "404.tmpl", gin.H{})
}

func main() {
	router := gin.Default()
	router.GET("/", WelcomeIndex)
	router.GET("/about-us", AboutUs)
	router.GET("/login", Login)
	router.POST("/submit-login", SubmitLogin) //          <-- a POST to this action
	router.NoRoute(Custom404)
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/assets", "assets")
	port := 3000
	router.Run(fmt.Sprintf(":%d", port))
}
