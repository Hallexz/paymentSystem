package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"paymentSystem/src"
	"testing"

	pb "paymentSystem/proto"

	"github.com/gin-gonic/gin"
)

func TestConfirmPayment(t *testing.T) {
	// Создаем тестовый HTTP-сервер
	router := gin.Default()
	s := &src.PaymentService{}
	router.POST("/payment/confirm", func(c *gin.Context) {
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

	// Создаем тестовый HTTP-запрос
	paymentRequest := pb.PaymentRequest{
		// Заполняем поля запроса
	}
	jsonPayload, err := json.Marshal(paymentRequest)
	if err != nil {
		t.Errorf("Failed to marshal payment request: %v", err)
	}

	req, err := http.NewRequest("POST", "/payment/confirm", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
	}

	// Создаем записыватель для захвата ответа
	w := httptest.NewRecorder()

	// Отправляем запрос
	router.ServeHTTP(w, req)

	// Проверяем код ответа
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Проверяем содержимое ответа (если необходимо)
	// ...
}
