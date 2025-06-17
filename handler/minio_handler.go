package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/service"
)

type MinioHandler struct {
	minioService *service.MinioService
}

func NewMinioHandler(minioService *service.MinioService) *MinioHandler {
	return &MinioHandler{minioService: minioService}
}

func (h *MinioHandler) GetInfo(c *gin.Context) {
	c.Set("handler_result", h.minioService)

	//c.JSON(http.StatusOK, gin.H{"user":},
}

func (h *MinioHandler) GetBucketList(c *gin.Context) {
	c.Set("handler_result", h.minioService.GetBucketList())
}
