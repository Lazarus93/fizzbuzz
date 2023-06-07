package router

import (
	"fizzbuzz/internal/fizzbuzz"
	"fizzbuzz/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, fbService *fizzbuzz.Service) {
	h := handler.NewHandler(fbService)

	r.POST("/fizzbuzz", h.FizzBuzzHandler)
}
