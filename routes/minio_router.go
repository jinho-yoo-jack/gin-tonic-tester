package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/handler"
)

type MinioRouters struct {
	minioHandler *handler.MinioHandler
}

func NewMinioRouters(minioHandler *handler.MinioHandler) *MinioRouters {
	return &MinioRouters{minioHandler: minioHandler}
}

func (mr *MinioRouters) RegisterRoutes(router *gin.Engine, authorize gin.HandlerFunc) {
	api := router.Group("/api/v1/minio")
	{
		api.GET("/info", mr.minioHandler.GetInfo)
		api.GET("/bucket/list", mr.minioHandler.GetBucketList)
	}
}
