package handlers

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	r := router.Group("/api")
	RegisterHello(r)
	RegisterPay(r)
}
