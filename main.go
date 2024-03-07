package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.StaticFile("/letter", "./static/Letter/index.html")
	r.Static("/css", "./static/Letter/")
	r.Run()
}
