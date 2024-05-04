package api

import (
	"context"
	"net/http"
	"paymentSystem/src"

	"github.com/gin-gonic/gin"

	pb "paymentSystem/proto"
)

func respon() {
	r := gin.Default()
	s := &src.PaymentService{}

	r.POST("/payment/process", func(c *gin.Context) {
		var req pb.PaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := s.ProcessPayment(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.Run()
}
