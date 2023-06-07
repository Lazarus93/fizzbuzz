package handler

import (
	"fizzbuzz/internal/fizzbuzz"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	fbService *fizzbuzz.Service
}

type Request struct {
	String1 string `json:"string1" swag:"example,fizz"`
	String2 string `json:"string2" swag:"example,buzz"`
	Int1    int    `json:"int1" swag:"example,3"`
	Int2    int    `json:"int2" swag:"example,5"`
	Limit   int    `json:"limit" swag:"example,15"`
}

func NewHandler(fbService *fizzbuzz.Service) *Handler {
	return &Handler{fbService: fbService}
}

// FizzBuzzHandler godoc
// @Summary Generate a FizzBuzz sequence
// @Description Generate a FizzBuzz sequence based on the provided parameters
// @ID fizzbuzz-handler
// @Accept  json
// @Produce  json
// @Param body body handler.Request true "FizzBuzz parameters"
// @Success 200 {array} string
// @Router /fizzbuzz [post]
func (h *Handler) FizzBuzzHandler(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seq, err := h.fbService.GenerateSequence(req.String1, req.String2, req.Int1, req.Int2, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": seq})
}
